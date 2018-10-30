package http

import (
	"ecards/structs"

	"github.com/astaxie/beego/context"
)

func SendOutput(
	c *context.Context,
	jsonBody interface{},
	err error,
) {
	if err != nil {
		c.Output.SetStatus(400)
	}

	var respData structs.RespData
	respData.RespBody = jsonBody
	respData.Error = err

	var t interface{} = respData

	errEnd := c.Output.JSON(t, true, false)
	if errEnd != nil {
		panic("Error Output JSON Level Middleware")
	}
}
