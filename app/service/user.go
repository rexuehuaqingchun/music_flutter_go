package service

import (
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

	constant.DetectionUserRoles(r.Type)

	listModel := m.Page(r.Page, r.Limit)
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.UserServiceGetListRes{
		Page:  r.Page,
		Size:  r.Limit,
		Total: total,
	}

	if err := listModel.ScanList(&getListRes.List, "Content"); err != nil {
		return nil, err
	}
	return getListRes, nil
}
