package service

import (
	"github.com/gogf/gf/frame/g"
	"music/app/dao"
	"music/app/model"
	"music/app/model/define"
	"music/constant"
)

type userService struct {
}

var User = userService{}

func (s userService) List(r *model.UserApiListReq) (*define.UserServiceGetListRes, error) {

	m := dao.User.Fields(model.UserListItem{})

	listModel := m.Page(r.Page, r.Limit)

	if constant.DetectionUserRoles(r.Type) {
		listModel = listModel.Where("type", r.Type)
	}

	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.UserServiceGetListRes{
		Page:  r.Page,
		Size:  r.Limit,
		Total: total,
	}

	if err := listModel.Structs(&getListRes.List); err != nil {
		return nil, err
	}
	return getListRes, nil
}

func (s userService) Info(r *model.UserApiInfoReq) (*define.UserInfo, error) {
	infoModel, err := dao.User.Where(g.Map{"id": r.Uid}).One()
	infoRes := &define.UserInfo{}
	infoModel.Struct(&infoRes.Info)
	if err != nil {
		return infoRes, err
	}
	return infoRes, err
}
