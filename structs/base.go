package structs

type (
	// ReqData Base All Request
	ReqData struct {
		ReqBody interface{} `json:"rqBody"`
	}

	// RespData All Response
	RespData struct {
		ResponseBody interface{}       `json:"rsBody"`
		Error        []FilterErrorCode `json:"error"`
	}

	// FilterErrorCode For error code
	FilterErrorCode struct {
		Code    string `json:"errorCode"`
		Message string `json:"errorDesc"`
	}
)
