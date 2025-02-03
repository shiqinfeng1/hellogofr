package route

import (
	"context"

	"hellogofr/pkg/openapi"
	"hellogofr/pkg/response"

	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/util/gmeta"
	"gofr.dev/pkg/gofr"
)

func BindOpenapiHandler(app *gofr.App) {
	app.GET("/openapi.json", func(ctx *gofr.Context) (interface{}, error) {
		return response.File(openapi.Bytes()), nil
	})
}

func BindHandler[Req any, Res any](app *gofr.App, handler func(*gofr.Context, *Req) (*Res, error)) {
	var req Req
	method := gmeta.Get(req, "method").String()
	path := gmeta.Get(req, "path").String()

	// 绑定uri和处理函数
	switch method {
	case "post":
		app.POST(path, gofrApiWrap(handler))
	case "get":
		app.GET(path, gofrApiWrap(handler))
	default:
		return
	}
	// 生成openapi接口描述
	openApiAdd(path, method, func(ctx context.Context, req *Req) (res *Res, err error) {
		return
	})
}

func gofrApiWrap[Req any, Res any](do func(*gofr.Context, *Req) (*Res, error)) func(*gofr.Context) (interface{}, error) {
	return func(ctx *gofr.Context) (interface{}, error) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			return nil, err
		}
		res, err := do(ctx, &req)
		if err != nil {
			return response.Error(1, err.Error()), nil
		}
		return response.Data(res), nil
	}
}

func openApiAdd[Req any, Res any](path, method string, f func(context.Context, *Req) (*Res, error)) {
	input := goai.AddInput{
		Path:   path,
		Method: method,
		Object: f,
	}
	err := openapi.Add(input)
	if err != nil {
		return
	}
}
