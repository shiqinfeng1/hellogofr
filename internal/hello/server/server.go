package server

import (
	"hellogofr/internal/hello/server/restful"
	"hellogofr/pkg/route"

	"gofr.dev/pkg/gofr"
)

func New() *gofr.App {
	app := gofr.New()

	// register api for openapi
	route.BindOpenapiHandler(app)

	// register http api and initialize openapi
	route.BindHandler(app, restful.GetConfig)

	return app
}
