package service

import (
	"math/rand"
	"music/app/dao"
	"music/app/model"
)

var Recommend = recommendService{}

type recommendService struct {
}

func (s recommendService) List(r *model.RecommendApiListReq) []model.QueryRecommendListItem {

	randModels := [3]string{"user", "video", "article"}
	models := getRandModels(r.Limit, randModels)
	result := make([]model.QueryRecommendListItem, r.Limit)
	for i := 0; i < len(models); i++ {
		key := models[i]
		switch key {
		case randModels[0]:
			users := getUsers(r.Page, 1)
			result[i].Type = randModels[0]
			result[i].UserList = users[0]
			break
		case randModels[1]:
			videos := getVideos(r.Page, 1)
			result[i].Type = randModels[1]
			result[i].VideoList = (*videos)[0]
			break
		case randModels[2]:
			articles := getArticle(r.Page, 1)
			result[i].Type = randModels[2]
			result[i].ArticleList = (*articles)[0]
			break
		}
	}

	return result

}

func getRandModels(limit int, models [3]string) []string {
	randModels := make([]string, limit)
	for i := 0; i < limit; i++ {
		var x = rand.Intn(len(models))
		randModels[i] = models[x]
	}
	return randModels
}

func getUsers(page int, limit int) []model.UserListItem {
	m := dao.User.Fields(model.UserListItem{})
	listModel := m.Page(page, limit)
	var resData []model.UserListItem
	listModel.Structs(&resData)
	return resData
}

func getVideos(page int, limit int) *[]model.QueryVideoListItem {
	var queryVideo = &[]model.QueryVideoListItem{}
	dao.Video.With(model.QueryUser{}).Page(page, limit).Scan(queryVideo)
	return queryVideo
}

func getArticle(page int, limit int) *[]model.QueryArticleListItem {
	var queryArticle = &[]model.QueryArticleListItem{}
	dao.Article.With(model.QueryUser{}).Page(page, limit).Scan(queryArticle)
	return queryArticle
}
