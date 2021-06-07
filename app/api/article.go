package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/app/service"
	"music/library/response"
)

var ArticleApi = articleApi{}

type articleApi struct {
}

func (v *articleApi) List(r *ghttp.Request) {
	var (
		reqData *model.ArticleApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.Article.List(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", list)
	}
}

func (*articleApi) Info(r *ghttp.Request) {
	var (
		reqData *model.ArticleApiInfoReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if info, err := service.Article.Info(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", info)
	}
}
