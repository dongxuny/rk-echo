// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package rkechopanic

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var defaultMiddlewareFunc = func(context echo.Context) error {
	return nil
}

func newCtx() echo.Context {
	return echo.New().NewContext(
		httptest.NewRequest(http.MethodGet, "/ut-path", nil),
		httptest.NewRecorder())
}

func TestInterceptor(t *testing.T) {
	defer assertNotPanic(t)

	handler := Interceptor(
		WithEntryNameAndType("ut-entry", "ut-type"))
	ctx := newCtx()

	f := handler(defaultMiddlewareFunc)

	assert.Nil(t, f(ctx))
}

func assertNotPanic(t *testing.T) {
	if r := recover(); r != nil {
		// Expect panic to be called with non nil error
		assert.True(t, false)
	} else {
		// This should never be called in case of a bug
		assert.True(t, true)
	}
}
