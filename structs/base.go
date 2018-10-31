package structs

type (
	// ReqData All Request
	ReqData struct {
		ReqBody interface{} `json:"rqBody"`
	}

	// RespData All Response
	RespData struct {
		ResponseBody interface{}       `json:"rsBody"`
		Error        []FilterErrorCode `json:"error"`
	}

	// FilterErrorCode ...
	FilterErrorCode struct {
		Code    string `json:"errorCode"`
		Message string `json:"errorDesc"`
	}
)
