package helper

import (
	"ecards/structs"

	"github.com/astaxie/beego/context"

	js "github.com/json-iterator/go"
)

var JS js.API

//GetRqBody ...
func GetRqBody(ctx *context.Context,
	errCode *[]structs.TypeError) []byte {
	var reqData structs.ReqData

	err1 := JS.Unmarshal(ctx.Input.RequestBody, &reqData)
	CheckErr("Error Unmarshal Get RqBody", err1)

	bodyByte, err2 := JS.Marshal(reqData.ReqBody)
	CheckErr("Error Unmarshal Get RqBody2", err2)

	if err1 != nil || err2 != nil {
		structs.ErrorCode.RequestMalformed.String(errCode)
	}

	return bodyByte

}
