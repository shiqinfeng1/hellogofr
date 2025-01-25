package route

import (
	"context"
	"net/http"

	"hellogofr/pkg/openapi"

	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/util/gmeta"
	"gofr.dev/pkg/gofr"
)

func BindHandler[Req any, Res any](app *gofr.App, handler func(ctx *gofr.Context) (interface{}, error)) {
	var req Req
	method := gmeta.Get(req, "method").String()
	path := gmeta.Get(req, "path").String()

	switch method {
	case "post":
		app.POST(path, handler)
		bindOpenApi[Req, Res]()
	case "get":
		app.GET(path, handler)
		bindOpenApi[Req, Res]()
	default:
		return
	}
}

func bindOpenApi[Req any, Res any]() {
	var req Req
	err := openapi.Add(goai.AddInput{
		Path:   gmeta.Get(req, "path").String(),
		Method: http.MethodPost,
		Object: func(ctx context.Context, req *Req) (res *Res, err error) { // 为了适配goai包生成openapi.json文件
			return
		},
	})
	if err != nil {
		return
	}
}
