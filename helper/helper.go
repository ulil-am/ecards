package helper

import (
	"ecards/structs"

	"github.com/astaxie/beego/context"

	js "github.com/json-iterator/go"
)

//GetRqBody ...
func GetRqBody(ctx *context.Context, err *error) []byte {
	var reqData structs.ReqData

	err1 := js.Unmarshal(ctx.Input.RequestBody, &reqData)
	CheckErr("Error Unmarshal Get RqBody", err1)

	bodyByte, err2 := js.Marshal(reqData.ReqBody)
	CheckErr("Error Unmarshal Get RqBody2", err2)

	if err1 != nil {
		*err = err1
	}

	if err2 != nil {
		*err = err2
	}

	return bodyByte

}
