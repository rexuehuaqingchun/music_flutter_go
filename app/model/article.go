// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
	"music/app/model/internal"
)

// Article is the golang structure for table article.
type Article internal.Article

// Fill with you ideas below.

type ArticleApiListReq struct {
	Page  int `v:"required|min:1#页码不能为空|页面最小为1"`
	Limit int `v:"required|min:5|max:10#必须指定每页的数据量|每页数据最少不少于5条，最多不大于50|每页数据最少不少于5条，最多不大于50"`
}

type ArticleApiInfoReq struct {
	ArticleId uint `v:"required#文章id不能为空"`
}

type ArticleListItem struct {
	Id           uint   `json:"id"`
	Uid          uint   `json:"uid"`
	coverUrlList string `json:"cover_url_list"`
	Title        string `json:"title"`
	CommentCount string `json:"comment_count"`
	thumbUpCount uint   `json:"thumb_up_count"`
	readCount    uint   `json:"read_count"`
}
