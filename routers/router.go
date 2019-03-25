package routers

import (
	"export_excel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/sql",
				beego.NSRouter("/Excel",&controllers.ExcelController{},"post:Excel"),
			),
		)
	beego.AddNamespace(ns)
}
