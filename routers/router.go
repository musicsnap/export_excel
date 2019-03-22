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
				beego.NSInclude(
					&controllers.MainController{},
				),
			),
			beego.NSNamespace("/sql2",
				beego.NSRouter("/get",&controllers.MainController{},"get:Sql2"),
			),
		)
	beego.AddNamespace(ns)
}
