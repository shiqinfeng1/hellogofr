package restful

import (
	v1 "hellogofr/api/restful/hello/v1"

	"gofr.dev/pkg/gofr"
)

func DoHello(ctx *gofr.Context) (interface{}, error) {
	var req v1.GetConfigReq
	ctx.Bind(&req)
	return "Hello World!", nil
}
