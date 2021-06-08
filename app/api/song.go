package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/app/service"
	"music/library/response"
)

var Song = songApi{}

type songApi struct {
}

func (*songApi) List(r *ghttp.Request) {
	var (
		reqData *model.SongApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if list, err := service.Song.List(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", list)
	}
}

func (*songApi) Info(r *ghttp.Request) {
	var (
		reqData *model.SongApiInfoReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if info, err := service.Song.Info(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", info)
	}
}
