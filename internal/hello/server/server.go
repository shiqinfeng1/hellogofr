// Copyright @2025-2028 <SieYuan> . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/shiqinfeng1/hellogofr/internal/hello/server/restful"
	"github.com/shiqinfeng1/hellogofr/pkg/route"

	"gofr.dev/pkg/gofr"
)

func New() *gofr.App {
	app := gofr.New()

	// app.UseMiddleware(response.UpdatePostStatusCode())

	// register api for openapi
	route.BindOpenapiHandler(app)
	// register http api and initialize openapi
	route.BindHandler(app, restful.GetConfig)

	return app
}
