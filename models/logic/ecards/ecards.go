package ecards

import (
	structAPI "ecards/structs/api/http"
	structDb "ecards/structs/db"

	"github.com/astaxie/beego"
)

//InsertECards ...
func InsertECards(
	reqCreateEcards structAPI.ReqCreateECards,
	err *error,
) {
	var doc structDb.ECards

	doc.CardNumber = reqCreateEcards.CardNumber
	doc.Company = reqCreateEcards.Company
	doc.ExpiryDate = reqCreateEcards.ExpiryDate
	doc.Name = reqCreateEcards.Name

	*err = DBEcards.InsertECards(&doc)

	if err != nil {
		beego.Warning("Error insert data ecards", *err)
	}

	return
}
