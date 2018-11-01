package ecards

import (
	"ecards/helper"
	"ecards/structs"
	structAPI "ecards/structs/api/http"
	structDb "ecards/structs/db"

	"github.com/astaxie/beego"
)

//InsertECards ...
func InsertECards(
	reqCreateEcards structAPI.ReqCreateECards,
	errCode *[]structs.TypeError,
) {
	var (
		doc    structDb.ECards
		nmFunc = "Insert new cards"
	)

	doc.CardNumber = reqCreateEcards.CardNumber
	doc.Company = reqCreateEcards.Company
	doc.ExpiryDate = reqCreateEcards.ExpiryDate
	doc.Name = reqCreateEcards.Name
	beego.Debug("Logic Ecards ====> ", doc)
	err := DBEcards.InsertECards(&doc)
	beego.Debug(err)
	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
		helper.CheckErr(nmFunc+" ", err)
	}

	return
}
