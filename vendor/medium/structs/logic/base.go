package logic

import (
	"medium/structs"
	apiStructs "medium/structs/api"
)

type (
	// ContextStruct - Struct for contain all data related to context and put it in argument
	ContextStruct struct {
		HeaderAll    string
		Header       structs.ReqHTTPHeader
		HeaderTracer apiStructs.HeaderTracer
		JobID        string
	}
)
