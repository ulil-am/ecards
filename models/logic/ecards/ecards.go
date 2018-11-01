package ecards

import (
	"ecards/helper"
	"ecards/structs"
	structAPI "ecards/structs/api/http"
	structDb "ecards/structs/db"
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

	err := DBEcards.InsertECards(&doc)

	if err != nil {
		structs.ErrorCode.DatabaseError.String(errCode, err.Error(), nmFunc)
		helper.CheckErr(nmFunc+" ", err)
	}

	return
}
