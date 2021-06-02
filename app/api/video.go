package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Video = videoApi{}

type videoApi struct {
}

func (*videoApi) Index(r *ghttp.Request) {
	r.Response.Writeln("video list")
}
