package service

import (
	"github.com/gogf/gf/frame/g"
	"music/app/dao"
	"music/app/model"
	"music/app/model/define"
)

var Article = articleService{}

type articleService struct {
}

func (s articleService) List(r *model.ArticleApiListReq) (*define.ArticleServiceGetListRes, error) {

	m := dao.Article.Fields(model.Article{})

	listModel := m.Page(r.Page, r.Limit)

	total, err := m.Fields("*").Count()
	if err != nil {
		return nil, err
	}
	getListRes := &define.ArticleServiceGetListRes{
		Page:  r.Page,
		Size:  r.Limit,
		Total: total,
	}

	if err := listModel.Structs(&getListRes.List); err != nil {
		return nil, err
	}
	return getListRes, nil
}

func (s articleService) Info(r *model.ArticleApiInfoReq) (*define.ArticleInfo, error) {

	infoModel, err := dao.Article.Where(g.Map{"id": r.ArticleId}).One()
	infoRes := &define.ArticleInfo{}
	infoModel.Struct(&infoRes.Info)
	if err != nil {
		return infoRes, err
	}
	return infoRes, err
}
