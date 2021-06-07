package api

import (
	"github.com/gogf/gf/net/ghttp"
	"music/app/model"
	"music/app/service"
	"music/constant"
	"music/library/response"
)

var UserApi = userApi{}

type userApi struct{}

func (*userApi) List(r *ghttp.Request) {
	var (
		reqData *model.UserApiListReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if reqData.Limit == 0 {
		reqData.Limit = constant.DefaultPageCount
	}

	if list, err := service.User.List(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", list)
	}

}

func (*userApi) Info(r *ghttp.Request) {
	var (
		reqData *model.UserApiInfoReq
	)
	if err := r.Parse(&reqData); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if info, err := service.User.Info(reqData); err != nil {
		response.JsonExit(r, 0, err.Error())
	} else {
		response.JsonExit(r, 0, "", info)
	}
}
