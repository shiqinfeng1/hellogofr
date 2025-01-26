package server

import (
	v1 "hellogofr/api/restful/hello/v1"
	"hellogofr/internal/hello/server/restful"
	"hellogofr/pkg/openapi"
	"hellogofr/pkg/response"
	"hellogofr/pkg/route"

	"gofr.dev/pkg/gofr"
)

func New() *gofr.App {
	app := gofr.New()

	// register openapi
	app.GET("/openapi.json", func(ctx *gofr.Context) (interface{}, error) {
		return response.Json(openapi.Bytes()), nil
	})

	// register http api
	route.BindHandler[v1.GetConfigReq, v1.GetConfigRes](app, restful.DoHello)

	return app
}
