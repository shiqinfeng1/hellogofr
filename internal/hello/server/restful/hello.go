// Copyright @2025-2028 <SieYuan> . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	v1 "github.com/shiqinfeng1/hellogofr/api/restful/hello/v1"

	"gofr.dev/pkg/gofr"
)

func GetConfig(ctx *gofr.Context, req *v1.GetConfigReq) (*v1.GetConfigRes, error) {
	// todo anything
	return &v1.GetConfigRes{}, nil
}
