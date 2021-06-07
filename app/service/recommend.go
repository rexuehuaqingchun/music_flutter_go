package service

import (
	"github.com/gogf/gf/frame/g"
	"music/app/dao"
	"music/app/model"
)

var Recommend = recommendService{}

type recommendService struct {
}

func (s recommendService) List(r *model.RecommendApiListReq) {
	getVideos(r.Page, r.Limit)
}

func getData(page int, limit int) {

	for i := 0; i < page; i++ {

	}
}

func getUsers(page int, limit int) []*model.UserListItem {
	m := dao.User.Fields(model.UserListItem{})
	listModel := m.Page(page, limit)
	var resData []*model.UserListItem
	listModel.Structs(&resData)
	return resData
}

func getVideos(page int, limit int) []*model.VideoListItem {

	m := dao.User.Fields(model.VideoListItem{})
	listModel := m.Page(page, limit)
	var resData []*model.VideoListItem
	listModel.Structs(&resData)

	for i := 0; i < len(resData); i++ {
		item := resData[i]
		videoId := item.Id
		var queryUser = model.QueryUser{}
		g.DB().Table("video").With(model.QueryUser{}).Where("id", videoId).Scan(queryUser)

	}

	return resData
}

func getArticle(page int, limit int) []*model.ArticleListItem {
	m := dao.User.Fields(model.ArticleListItem{})
	listModel := m.Page(page, limit)
	var resData []*model.ArticleListItem
	listModel.Structs(&resData)
	return resData
}
