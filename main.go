package main

import (
	"ecards/helper/constant"
	"ecards/helper/timetn"
	"ecards/routers/grpc"
	_ "ecards/routers/http"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	setup()

	if constant.GOENV != "" && constant.GOAPP == beego.BConfig.AppName {
		time.AfterFunc(1*time.Second, initialData)

		beego.BConfig.Listen.Graceful = true
		if beego.BConfig.RunMode == "dev" {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.Listen.EnableAdmin = true
			beego.BConfig.Listen.AdminAddr = "localhost"

			adminPort, _ := strconv.Atoi("2" + strconv.Itoa(beego.BConfig.Listen.HTTPPort))
			beego.BConfig.Listen.AdminPort = adminPort
			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		} else {
			// beego.SetLevel(beego.LevelInformational)
			beego.BConfig.RecoverPanic = true
			beego.BConfig.Listen.ServerTimeOut = 680
		}

		beego.Run()

	} else {
		beego.Error("SETUP GOENV && GOAPP FIRST")
	}
}

func setup() {
	logs.Async()

	logFile := logFile()
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"`+logFile+`",
		"level":7,"maxlines":20000,"maxsize":0,"daily":true,"maxdays":10}`)
	if err != nil {
		panic(err)
	}

	// db.RegisterPGSQL()
}

func logFile() string {
	hostname, _ := os.Hostname()
	now := timetn.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hours := now.Hour()

	strTime := strconv.Itoa(year) + `-` +
		strconv.Itoa(month) + `-` +
		strconv.Itoa(day) + `-` +
		strconv.Itoa(hours)

	pathLog := "logs/" + hostname + "_" + strTime +
		"_" + constant.GOAPP + ".log"
	return pathLog
}

func initialData() {
	grpc.CreateGrpcServer("")
	go initPprof()
}

func initPprof() {
	portHTTP := "1" + strconv.Itoa(beego.BConfig.Listen.HTTPPort)
	http.ListenAndServe(":"+portHTTP, nil)
}
