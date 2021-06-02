package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/library/response"
)

var Video = videoApi{}

type videoApi struct {
}

func (*videoApi) Index(r *ghttp.Request) {
	var (
		reqData *model.VideoApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
}
