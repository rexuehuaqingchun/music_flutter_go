package service

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
)

type middlewareService struct{}

var Middleware = middlewareService{}

func (s *middlewareService) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)

	r.Middleware.Next()
}

// 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
