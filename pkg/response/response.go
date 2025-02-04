// Copyright @2025-2028 <SieYuan> . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package response

import (
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

// UpdatePostStatusCode 把gofr框架默认post请求返回201更新为200
// func UpdatePostStatusCode() func(inner http.Handler) http.Handler {
// 	return func(inner http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			inner.ServeHTTP(w, r)
// 			if r.Method == http.MethodPost {
// 				w.WriteHeader(http.StatusOK)
// 			}
// 		})
// 	}
// }
