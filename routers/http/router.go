package routers

import (
	ctrl "ecards/controllers/http"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	Router()
}

func Router() {
	// beego.ErrorHandler("404", pageNotFound)
	ns := beego.NewNamespace("/ecards/v1",
		// beego.NSNamespace("/ecards",
		// 	beego.NSInclude(
		// 		&ctrl.EcardsController{},
		// 	),
		// ),
		beego.NSRouter("/ecards",
			&ctrl.EcardsController{},
			"post:Post"),
	)

	beego.AddNamespace(ns)
}
