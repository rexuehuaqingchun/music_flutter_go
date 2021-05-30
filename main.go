package main

import (
	_ "music/boot"
	_ "music/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
