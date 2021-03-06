// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
	"music/app/model/internal"
)

// User is the golang structure for table user.
type User internal.User

// Fill with you ideas below.
type UserApiListReq struct {
	Page  int `v:"required|min:1#页码不能为空|页面最小为1"`
	Limit int `v:"required|min:5|max:10#必须指定每页的数据量|每页数据最少不少于5条，最多不大于50|每页数据最少不少于5条，最多不大于50"`
	Type  string
}

type UserApiInfoReq struct {
	Uid uint `v:"required#用户id不能为空"`
}

type UserListItem struct {
	Id              uint   `json:"id"`
	CoverPictureUrl string `json:"coverPictureUrl"`
	Nickname        string `json:"nickname"`
	Type            string `json:"type"`
	MusicCount      uint   `json:"musicCount"`
	MusicPlayCount  uint   `json:"musicPlayCount"`
}
