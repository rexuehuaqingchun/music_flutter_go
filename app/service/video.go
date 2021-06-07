package service

import (
	"github.com/gogf/gf/frame/g"
	"music/app/dao"
	"music/app/model"
	"music/app/model/define"
)

type videoService struct {
}

var Video = videoService{}

func (s videoService) List(r *model.VideoApiListReq) (*define.VideoServiceGetListRes, error) {

	m := dao.Video.Fields(model.Video{})

	listModel := m.Page(r.Page, r.Limit)

	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.VideoServiceGetListRes{
		Page:  r.Page,
		Size:  r.Limit,
		Total: total,
	}

	if err := listModel.Structs(&getListRes.List); err != nil {
		return nil, err
	}
	return getListRes, nil
}

func (s videoService) Info(r *model.VideoApiInfoReq) (*define.VideoInfo, error) {
	infoModel, err := dao.Video.Where(g.Map{"id": r.VideoId}).One()
	infoRes := &define.VideoInfo{}
	infoModel.Struct(&infoRes.Info)
	if err != nil {
		return infoRes, err
	}
	return infoRes, err
}
