package model

import "github.com/gogf/gf/util/gmeta"

type RecommendApiListReq struct {
	Page  int `v:"required|min:1#页码不能为空|页面最小为1"`
	Limit int `v:"required|min:5|max:10#必须指定每页的数据量|每页数据最少不少于5条，最多不大于50|每页数据最少不少于5条，最多不大于50"`
}

type QueryUser struct {
	gmeta.Meta      `orm:"table:user"`
	Id              uint   `json:"id"`
	CoverPictureUrl string `json:"cover_picture_url"`
	Nickname        string `json:"nickname"`
	Type            string `json:"type"`
	MusicCount      uint   `json:"music_count"`
	MusicPlayCount  uint   `json:"music_play_count"`
}

type QueryVideoListItem struct {
	gmeta.Meta      `orm:"table:video"`
	Id              uint       `json:"id"`
	UserId          uint       `json:"user_id"`
	CoverPictureUrl string     `json:"cover_picture_url"`
	VideoUrl        string     `json:"video_url"`
	Title           string     `json:"title"`
	Intro           string     `json:"intro"`
	CommentCount    uint       `json:"comment_count"`
	ThumbUpCount    uint       `json:"thumb_up_count"`
	ReadCount       uint       `json:"read_count"`
	ShareCount      uint       `json:"share_count"`
	ContentSeconds  uint       `json:"content_seconds"`
	User            *QueryUser `orm:"with:id=UserId"`
}

type QueryArticleListItem struct {
	gmeta.Meta   `orm:"table:article"`
	Id           uint       `json:"id"`
	UserId       uint       `json:"user_id"`
	coverUrlList string     `json:"cover_url_list"`
	Title        string     `json:"title"`
	CommentCount string     `json:"comment_count"`
	thumbUpCount uint       `json:"thumb_up_count"`
	readCount    uint       `json:"read_count"`
	User         *QueryUser `orm:"with:id=UserId"`
}

type QueryRecommendListItem struct {
	Type        string               `json:"type"`
	UserList    UserListItem         `json:"user_entry"`
	VideoList   QueryVideoListItem   `json:"video_list"`
	ArticleList QueryArticleListItem `json:"article_list"`
}
