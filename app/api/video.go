package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/app/service"
	"music/library/response"
)

var Video = videoApi{}

type videoApi struct {
}

func (v *videoApi) List(r *ghttp.Request) {
	var (
		reqData *model.VideoApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.Video.List(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", list)
	}
}

func (*videoApi) Info(r *ghttp.Request) {
	var (
		reqData *model.VideoApiInfoReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if info, err := service.Video.Info(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", info)
	}
}
