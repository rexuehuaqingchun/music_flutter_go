package service

import (
	"errors"
	"fmt"
	"music/app/model"
)

type videoService struct {
}

var Video = videoService{}

func (v videoService) List(r *model.VideoApiListReq) error {
	if r.Page < 0 {
		return errors.New(fmt.Sprintf("页码 %s 不能小于0", r.Page))
	}
	return nil
}
