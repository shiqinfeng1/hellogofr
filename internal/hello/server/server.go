package server

import (
	v1 "hellogofr/api/restful/hello/v1"
	"hellogofr/internal/hello/server/restful"
	"hellogofr/pkg/route"

	"gofr.dev/pkg/gofr"
)

func New() *gofr.App {
	app := gofr.New()

	route.BindHandler[v1.GetConfigReq, v1.GetConfigRes](app, restful.DoHello)

	return app
}
