// Copyright @2025-2028 <SieYuan> . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package response

import (
	"bytes"
	"net/http"

	"gofr.dev/pkg/gofr/http/response"
)

type resp struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail string `json:"detail,omitempty"`
	Data   any    `json:"data"`
}

func File(b []byte) response.File {
	return response.File{Content: b, ContentType: "application/json"}
}

func Data(data any) response.Raw {
	return response.Raw{
		Data: resp{
			Code: 0,
			Msg:  "",
			Data: data,
		},
	}
}

func Error(code int, msg, detail string) response.Raw {
	return response.Raw{
		Data: resp{
			Code:   code,
			Msg:    msg,
			Detail: detail,
			Data:   "",
		},
	}
}

// responseWriter 自定义响应写入器，用于捕获响应数据
type responseWriter struct {
	http.ResponseWriter
	buf bytes.Buffer
}

func (lrw *responseWriter) Write(b []byte) (int, error) {
	// 将响应数据写入缓冲区
	return lrw.buf.Write(b)
}

// Logging is a middleware which logs response status and time in milliseconds along with other data.
func BuildResponse() func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			srw := &responseWriter{ResponseWriter: w}

			defer func(res *responseWriter, req *http.Request) {
			}(srw, r)

			inner.ServeHTTP(srw, r)
		})
	}
}
