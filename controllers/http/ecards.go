package http

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"ecards/helper"
	logicECards "ecards/models/logic/ecards"
	structAPI "ecards/structs/api/http"
)

type EcardsController struct {
	beego.Controller
}

func (c *EcardsController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *EcardsController) Post() {
	var (
		err          error
		reqInterface structAPI.CreateECardsInteface
		req          structAPI.ReqCreateECards
	)
	rqBodyByte := helper.GetRqBody(c.Ctx, &err)
	if err != nil {
		SendOutput(c.Ctx, c.Data["json"], err)
		return
	}

	err2 := json.Unmarshal(rqBodyByte, &reqInterface)
	if err2 != nil {
		SendOutput(c.Ctx, c.Data["json"], err)
		return
	}

	// reqInterface.ValidateRequest(&req, &err)

	req.CardNumber = int(rqBodyByte[0])
	req.Company = string(rqBodyByte[1])
	req.ExpiryDate = string(rqBodyByte[2])
	req.Name = string(rqBodyByte[3])

	logicECards.InsertECards(req, &err)

}
