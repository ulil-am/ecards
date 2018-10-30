package main

import (
	_ "ecards/routers/http"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.Graceful = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
