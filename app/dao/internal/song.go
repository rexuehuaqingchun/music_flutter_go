// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SongDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SongDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns songColumns // Columns contains all the columns of Table that for convenient usage.
}

// SongColumns defines and stores column names for table song.
type songColumns struct {
	Id              string //
	UserId          string //
	CoverPictureUrl string //
	SongUrl         string //
	CnName          string //
	EnName          string //
	CommentCount    string //
	ThumbUpCount    string //
	ReadCount       string //
}

var (
	// Song is globally public accessible object for table song operations.
	Song = SongDao{
		M:     g.DB("default").Model("song").Safe(),
		DB:    g.DB("default"),
		Table: "song",
		Columns: songColumns{
			Id:              "id",
			UserId:          "userId",
			CoverPictureUrl: "coverPictureUrl",
			SongUrl:         "songUrl",
			CnName:          "cnName",
			EnName:          "enName",
			CommentCount:    "commentCount",
			ThumbUpCount:    "thumbUpCount",
			ReadCount:       "readCount",
		},
	}
)