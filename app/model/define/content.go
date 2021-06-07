package define

import "music/app/model"

type UserServiceGetListRes struct {
	List  []*model.UserListItem `json:"list"`  // 列表
	Page  int                   `json:"page"`  // 分页码
	Size  int                   `json:"size"`  // 分页数量
	Total int                   `json:"total"` // 数据总数
}

type UserInfo struct {
	Info *model.UserListItem `json:"info"`
}

type VideoServiceGetListRes struct {
	List  []*model.VideoListItem `json:"list"`  // 列表
	Page  int                    `json:"page"`  // 分页码
	Size  int                    `json:"size"`  // 分页数量
	Total int                    `json:"total"` // 数据总数
}

type VideoInfo struct {
	Info *model.VideoListItem `json:"info"`
}

type ArticleServiceGetListRes struct {
	List  []*model.ArticleListItem `json:"list"`  // 列表
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

type ArticleInfo struct {
	Info *model.ArticleListItem `json:"info"`
}
