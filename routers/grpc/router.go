// @APIVersion 1.0.0
// @Title DLOR Acquisition GRPC API Documentation
// @Description DLOR Acquisition collect and acquisite the customer data
package grpc

import (
	ctrl "ecards/controllers/grpc"
	"ecards/helper/constant"
	pb "ecards/proto"
	"ecards/structs"
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
		prefix + "/acquisition": ctrl.RPCtrlEcards,
		/*:ENDGRPC*/
	}
}
