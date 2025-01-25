package openapi

import (
	"sync"

	"github.com/gogf/gf/v2/net/goai"
)

var (
	once sync.Once
	oai  *goai.OpenApiV3
)

func init() {
	once.Do(func() {
		oai = goai.New()
		oai.Info.Title = "GoOpenAPI"
		oai.Info.Version = "v1.0"
		oai.Info.Description = "GoOpenAPI"
	})
}

func Add(in goai.AddInput) error {
	return oai.Add(in)
}

func String() string {
	return oai.String()
}
