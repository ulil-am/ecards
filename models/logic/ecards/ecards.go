package ecards

import (
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
	var doc structDb.ECards
	var err error
	doc.CardNumber = reqCreateEcards.CardNumber
	doc.Company = reqCreateEcards.Company
	doc.ExpiryDate = reqCreateEcards.ExpiryDate
	doc.Name = reqCreateEcards.Name

	err = DBEcards.InsertECards(&doc)

	if err != nil {
		beego.Warning("Error insert data ecards", err)
	}

	return
}
