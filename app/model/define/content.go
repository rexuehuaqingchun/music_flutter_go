package define

import "music/app/model"

type UserServiceGetListRes struct {
	List  []*model.UserListItem `json:"list"`  // 列表
	Page  int                   `json:"page"`  // 分页码
	Size  int                   `json:"size"`  // 分页数量
	Total int                   `json:"total"` // 数据总数
}
