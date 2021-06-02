package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Hello = userApi{}

type userApi struct{}

func (*userApi) List(r *ghttp.Request) {
	r.Response.Writeln("")
}
