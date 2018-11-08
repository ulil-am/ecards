package grpc

import (
	"ecards/helper"
	pb "ecards/proto"
	"ecards/structs"
	structRpc "ecards/structs/api/grpc"
	"encoding/json"
)

// RPCtrlEcards ...
func RPCtrlEcards(
	in *pb.DoReq,
	errRPCCode *structs.TypeGRPCError,
	body *[]byte,
) {
	var (
		req structRpc.ReqTest
		res structRpc.ResTest
	)

	err := json.Unmarshal(in.GetBody(), &req)
	if err != nil {
		helper.CheckErr("Failed unmarshal @RPCtrlEcards", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	res.ID = req.ID
	res.Res = "response"
	resBy, err := json.Marshal(res)
	if err != nil {
		helper.CheckErr("Failed marshal @RPCtrlEcards", err)
		structs.ErrorCode.UnexpectedError.String(&errRPCCode.Error)
		return
	}

	*body = resBy

}
