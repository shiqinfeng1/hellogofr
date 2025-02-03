package restful

import (
	v1 "hellogofr/api/restful/hello/v1"

	"gofr.dev/pkg/gofr"
)

func GetConfig(ctx *gofr.Context, req *v1.GetConfigReq) (*v1.GetConfigRes, error) {
	// todo anything
	return &v1.GetConfigRes{}, nil
}
