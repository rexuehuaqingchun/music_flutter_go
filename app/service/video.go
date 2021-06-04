package service

import (
	"music/app/dao"
	"music/app/model"
)

type videoService struct {
}

var Video = videoService{}

func (s videoService) List(r *model.VideoApiListReq) error {

	fields := dao.Video.Fields(model.VideoListItem{})
	if r.Type != "" {
		fields.Where(dao.Video.Columns.CommentCount)
	}
	return nil
}
