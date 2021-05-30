package model

import "github.com/gogf/gf/net/ghttp"

const (
	ContextKey = "ContextKey"
)

type Context struct {
	Session *ghttp.Session
}

type ContextUser struct {
	Id uint //用户id
}
