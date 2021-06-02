// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// Song is the golang structure for table song.
type Song struct {
	Id              int    `orm:"id,primary"      json:"id"`              //
	UserId          int    `orm:"userId"          json:"userId"`          //
	CoverPictureUrl string `orm:"coverPictureUrl" json:"coverPictureUrl"` //
	SongUrl         string `orm:"songUrl"         json:"songUrl"`         //
	CnName          string `orm:"cnName"          json:"cnName"`          //
	EnName          string `orm:"enName"          json:"enName"`          //
	CommentCount    int    `orm:"commentCount"    json:"commentCount"`    //
	ThumbUpCount    int    `orm:"thumbUpCount"    json:"thumbUpCount"`    //
	ReadCount       int    `orm:"readCount"       json:"readCount"`       //
}