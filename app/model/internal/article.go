// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Article is the golang structure for table article.
type Article struct {
	Id           int    `orm:"id,primary"   json:"id"`           //
	UserId       int    `orm:"userId"       json:"userId"`       //
	CoverUrlList string `orm:"coverUrlList" json:"coverUrlList"` // 多个英文逗号分割 最多四张
	Title        string `orm:"title"        json:"title"`        //
	CommentCount int    `orm:"commentCount" json:"commentCount"` //
	ThumbUpCount int    `orm:"thumbUpCount" json:"thumbUpCount"` //
	ReadCount    int    `orm:"readCount"    json:"readCount"`    //
}
