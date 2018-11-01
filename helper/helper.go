package helper

import (
	"ecards/structs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"

	js "github.com/json-iterator/go"
)

var JS js.API

// Bm ...
var Bm cache.Cache

func init() {
	bm, err := cache.NewCache("memory", `{"interval":0}`)
	Bm = bm
	if err != nil {
		beego.Warning("error set cache", err)
	}
	// JS = js.ConfigFastest
	JS = js.ConfigCompatibleWithStandardLibrary
}

//GetRqBody ...
func GetRqBody(ctx *context.Context,
	errCode *[]structs.TypeError) []byte {
	var reqData structs.ReqData

	err1 := JS.Unmarshal(ctx.Input.RequestBody, &reqData)
	beego.Debug(err1)
	CheckErr("Error Unmarshal Get RqBody12321", err1)

	bodyByte, err2 := JS.Marshal(reqData.ReqBody)
	beego.Debug(err2)
	CheckErr("Error Marshal Get RqBody2321321", err2)

	if err1 != nil || err2 != nil {
		structs.ErrorCode.RequestMalformed.String(errCode)
	}

	return bodyByte

}
