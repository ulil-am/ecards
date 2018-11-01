package grpc

import (
	ctrl "medium/controllers/grpc"
	"medium/helper/constant"
	pb "medium/proto"
	"medium/structs"
)

var (
	prefix = "/" + constant.GOAPP + "/" + constant.VERSION
)

type fnRouteRPC func(
	*pb.DoReq,
	*structs.TypeGRPCError,
	*[]byte,
)

var routeMap map[string]fnRouteRPC

func init() {
	Router()
}

func Router() {
	routeMap = map[string]fnRouteRPC{
		/*:STARTGRPC*/
		prefix + "/acquisition": ctrl.RPCtrlAcquisition,
		/*:ENDGRPC*/
	}
}
