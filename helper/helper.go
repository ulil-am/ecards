package helper

import (
	"ecards/structs"
	structRPC "ecards/structs/api/grpc"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"

	js "github.com/json-iterator/go"

	"ecards/helper/timetn"
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

// SetHeaderRPC ...
func SetHeaderRPC(befMSStr string, reqID string,
	errorHeader structs.TypeGRPCError) []byte {
	nowMs := timetn.Now().UnixNano() / int64(time.Millisecond)

	befMs, err := strconv.ParseInt(befMSStr, 10, 64)
	CheckErr("", err)
	afterMs := nowMs - befMs

	header := structRPC.TypeHeaderRPC{
		ReqID:       reqID,
		Date:        timetn.Now(),
		ContentType: "application/grpc",
		RoundTrip:   strconv.FormatInt(afterMs, 10),
		Error:       errorHeader,
	}

	jsonByte, err := json.Marshal(header)
	CheckErr("", err)

	return jsonByte

}
