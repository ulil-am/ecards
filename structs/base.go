package structs

type (
	// ReqData All Request
	ReqData struct {
		ReqBody interface{} `json:"rqBody"`
	}

	// RespData All Response
	RespData struct {
		RespBody interface{} `json:"rsBody"`
		Error    error       `json:"error"`
	}
)
