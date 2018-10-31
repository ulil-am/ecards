package http

import (
	"ecards/structs"

	"github.com/astaxie/beego/context"
)

// SendOutput ...
func SendOutput(
	c *context.Context,
	jsonBody interface{},
	arrErrType []structs.TypeError,
) {
	if len(arrErrType) > 0 {
		c.Output.SetStatus(400)
	}

	var (
		hasIndent   = true
		hasEncoding = false
	)

	var respData structs.RespData
	respData.ResponseBody = jsonBody

	filErrCode := make([]structs.FilterErrorCode, 0)
	if len(arrErrType) > 0 {
		var t interface{}
		respData.ResponseBody = t
		filErrCode = filterErrorCode(arrErrType)
	}
	respData.Error = filErrCode

	var t interface{} = respData

	err := c.Output.JSON(t, hasIndent, hasEncoding)
	if err != nil {
		panic("Error Output JSON Level Middleware")
	}
}

func filterErrorCode(arrErrType []structs.TypeError) []structs.FilterErrorCode {
	var filter []structs.FilterErrorCode
	for _, val := range arrErrType {
		filter = append(filter, structs.FilterErrorCode{
			Code:    val.Code,
			Message: val.Message,
		})
	}
	return filter
}
