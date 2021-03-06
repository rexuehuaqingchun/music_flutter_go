// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// VideoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type VideoDao struct {
	gmvc.M               // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB       // DB is the raw underlying database management object.
	Table   string       // Table is the table name of the DAO.
	Columns videoColumns // Columns contains all the columns of Table that for convenient usage.
}

// VideoColumns defines and stores column names for table video.
type videoColumns struct {
	Id              string //
	UserId          string //
	CoverPictureUrl string //
	VideoUrl        string //
	Title           string //
	Intro           string //
	CommentCount    string //
	ThumbUpCount    string //
	ReadCount       string //
	ShareCount      string //
	ContentSeconds  string //
}

var (
	// Video is globally public accessible object for table video operations.
	Video = VideoDao{
		M:     g.DB("default").Model("video").Safe(),
		DB:    g.DB("default"),
		Table: "video",
		Columns: videoColumns{
			Id:              "id",
			UserId:          "userId",
			CoverPictureUrl: "coverPictureUrl",
			VideoUrl:        "videoUrl",
			Title:           "title",
			Intro:           "intro",
			CommentCount:    "commentCount",
			ThumbUpCount:    "thumbUpCount",
			ReadCount:       "readCount",
			ShareCount:      "shareCount",
			ContentSeconds:  "contentSeconds",
		},
	}
)
