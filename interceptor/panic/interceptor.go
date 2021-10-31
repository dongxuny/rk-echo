// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

// Package rkechopanic is a middleware of echo framework for recovering from panic
package rkechopanic

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rookie-ninja/rk-common/error"
	"github.com/rookie-ninja/rk-echo/interceptor"
	"github.com/rookie-ninja/rk-echo/interceptor/context"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

// Interceptor returns a echo.MiddlewareFunc (middleware)
func Interceptor(opts ...Option) echo.MiddlewareFunc {
	set := newOptionSet(opts...)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(rkechointer.RpcEntryNameKey, set.EntryName)

			defer func() {
				if recv := recover(); recv != nil {
					var res *rkerror.ErrorResp

					if se, ok := recv.(*rkerror.ErrorResp); ok {
						res = se
					} else if re, ok := recv.(error); ok {
						res = rkerror.FromError(re)
					} else {
						res = rkerror.New(rkerror.WithMessage(fmt.Sprintf("%v", recv)))
					}

					rkechoctx.GetEvent(ctx).SetCounter("panic", 1)
					rkechoctx.GetEvent(ctx).AddErr(res.Err)
					rkechoctx.GetLogger(ctx).Error(fmt.Sprintf("panic occurs:\n%s", string(debug.Stack())), zap.Error(res.Err))

					ctx.JSON(http.StatusInternalServerError, res)
				}
			}()

			return next(ctx)
		}
	}
}
