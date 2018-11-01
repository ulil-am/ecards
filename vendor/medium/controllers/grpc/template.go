package grpc

import (
	"encoding/json"
	"medium/helper"
	pb "medium/proto"
	"medium/structs"
	rpcStructs "medium/structs/api/grpc"
)

// RPCtrlAcquisition - RPCtrlAcquisition Controllers
func RPCtrlAcquisition(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {

	var (
		req rpcStructs.ReqTest
		res rpcStructs.ResTest
	)

	err := json.Unmarshal(in.GetBody(), &req)
	if err != nil {
		helper.CheckErr("failed unmarshal @RPCtrlAcquisition", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	res.ID = req.ID
	res.Res = "response"
	resBy, err := json.Marshal(res)
	if err != nil {
		helper.CheckErr("failed marshal &RPCtrlAcquisition", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	*body = resBy
}
