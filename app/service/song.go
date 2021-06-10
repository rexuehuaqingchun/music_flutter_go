package service

import (
	"github.com/gogf/gf/frame/g"
	"music/app/dao"
	"music/app/model"
	"music/app/model/define"
)

type songService struct {
}

var Song = songService{}

func (s songService) List(r *model.SongApiListReq) (*define.SongServiceGetListRes, error) {

	m := dao.Song.Fields(model.Video{})
	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.SongServiceGetListRes{
		Page:  r.Page,
		Size:  r.Limit,
		Total: total,
	}
	dao.Song.With(model.QueryUser{}).Page(r.Page, r.Limit).Scan(&getListRes.List)
	return getListRes, nil
}

func (s songService) Info(r *model.SongApiInfoReq) (*define.SongInfo, error) {
	infoModel, err := dao.Song.Where(g.Map{"id": r.SongId}).One()
	infoRes := &define.SongInfo{}
	infoModel.Struct(&infoRes.Info)
	if err != nil {
		return infoRes, err
	}
	return infoRes, err
}
