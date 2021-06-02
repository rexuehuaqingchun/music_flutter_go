package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"music/app/api"
	"music/app/service"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware.Ctx)
		group.ALL("/hello", api.Hello)
		group.ALL("/video", api.Video)
	})
}
