package response

import (
	"gofr.dev/pkg/gofr/http/response"
)

type resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Json(b []byte) response.File {
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

func Error(code int, msg string) response.Raw {
	return response.Raw{
		Data: resp{
			Code: code,
			Msg:  msg,
			Data: "",
		},
	}
}
