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
		group.ALLMap(g.Map{
			"/user":          api.UserApi,
			"/user/:info":    api.UserApi.Info,
			"/video":         api.Video,
			"/video/:info":   api.Video.Info,
			"/article":       api.ArticleApi,
			"/article/:info": api.ArticleApi.Info,
			"/recommend":     api.Recommend,
		})
	})
}
