package http

import (
	"ecards/structs"

	"github.com/jannes-sa/customvalidator"
)

type (
	//ReqCreateECards Struct for insert new card request
	ReqCreateECards struct {
		CardNumber int    `json:"card_number"`
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		Company    string `json:"company"`
	}

	//ECards Ecard struct
	ECards struct {
		CardNumber int    `json:"card_number"`
		Name       string `json:"name"`
		ExpiryDate string `json:"expiry_date"`
		Company    string `json:"company"`
	}

	// CreateECardsInteface ...
	CreateECardsInteface struct {
		CardNumber interface{} `json:"card_number" type:"int"`
		Name       interface{} `json:"name" type:"string"`
		ExpiryDate interface{} `json:"expiry_date" type:"string"`
		Company    interface{} `json:"company" type:"string"`
	}

	//RespCreateECards Struct for insert new card request
	RespCreateECards struct {
		CardNumber int `json:"card_number"`
	}
)

// ValidateRequest ...
func (st *CreateECardsInteface) ValidateRequest(
	assignStruct *ReqCreateECards,
	errCode *[]structs.TypeError,
) {
	var paramsInterChange interface{} = *st
	codeError := customvalidator.Validate(
		paramsInterChange, assignStruct,
	)
	structs.GetCodeError(codeError, errCode)
}
