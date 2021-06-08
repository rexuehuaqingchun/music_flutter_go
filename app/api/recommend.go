package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/app/service"
	"music/library/response"
)

var Recommend = recommendApi{}

type recommendApi struct {
}

func (*recommendApi) List(r *ghttp.Request) {
	var (
		reqData *model.RecommendApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "", service.Recommend.List(reqData))
}
