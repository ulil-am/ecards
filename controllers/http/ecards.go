package http

import (
	"encoding/json"
	"reflect"

	"ecards/structs"

	"github.com/astaxie/beego"

	"ecards/helper"
	logicECards "ecards/models/logic/ecards"
	structAPI "ecards/structs/api/http"
)

type EcardsController struct {
	beego.Controller
}

// func (c *EcardsController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// }

// Post ...
// @Title Create
// @Description create Products
// @Param	body		body 	models.ecards	true		"body for Products content"
// @Success 201 {object} models.Products
// @Failure 403 body is empty
// @router / [post]
func (c *EcardsController) Post() {
	errCode := make([]structs.TypeError, 0)

	var (
		reqInterface structAPI.CreateECardsInteface
		req          structAPI.ReqCreateECards
	)
	rqBodyByte := helper.GetRqBody(c.Ctx, &errCode)
	if len(errCode) > 0 {
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	err := json.Unmarshal(rqBodyByte, &reqInterface)
	if err != nil {
		structs.ErrorCode.UnexpectedError.String(&errCode)
		SendOutput(c.Ctx, c.Data["json"], errCode)
		return
	}

	reqInterface.ValidateRequest(&req, &errCode)

	req.CardNumber = int(reflect.ValueOf(reqInterface.CardNumber).Int())
	req.Company = string(reflect.ValueOf(reqInterface.Company).String())
	req.ExpiryDate = string(reflect.ValueOf(reqInterface.ExpiryDate).String())
	req.Name = string(reflect.ValueOf(reqInterface.Name).String())

	// req.CardNumber = int(rqBodyByte[0])
	// req.Company = string(rqBodyByte[1])
	// req.ExpiryDate = string(rqBodyByte[2])
	// req.Name = string(rqBodyByte[3])

	logicECards.InsertECards(req, &errCode)

}
