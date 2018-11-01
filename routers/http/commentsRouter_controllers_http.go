package http

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

// init ...
func init() {

	beego.GlobalControllerRouter["ecards/controllers/http:EcardsController"] = append(beego.GlobalControllerRouter["ecards/controllers/http:EcardsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
