// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns userColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserColumns defines and stores column names for table user.
type userColumns struct {
	Id              string //
	CoverPictureUrl string //
	Nickname        string //
	Type            string // 角色 NORMAL_USER DQ_SINGER DQ_OFFICIAL_ACCOUNT
	MusicCount      string //
	MusicPlayCount  string //
}

var (
	// User is globally public accessible object for table user operations.
	User = UserDao{
		M:     g.DB("default").Model("user").Safe(),
		DB:    g.DB("default"),
		Table: "user",
		Columns: userColumns{
			Id:              "id",
			CoverPictureUrl: "coverPictureUrl",
			Nickname:        "nickname",
			Type:            "type",
			MusicCount:      "musicCount",
			MusicPlayCount:  "musicPlayCount",
		},
	}
)
