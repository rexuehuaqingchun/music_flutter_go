package api

import (
	"fmt"
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

	if reqData.Page < 0 {
		response.JsonExit(r, 0, fmt.Sprintf("页码 %v不能小于0", reqData.Page))
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
