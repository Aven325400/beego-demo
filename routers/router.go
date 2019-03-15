// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"server-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.SetStaticPath("/swagger", "swagger")
	beego.Info("success")
	//ns := beego.NewNamespace("/test",
	//	beego.NSNamespace("/login",
	//		beego.NSInclude(
	//			&controllers.LoginController{},
	//		),
	//	),
	//)
	//beego.AddNamespace(ns)
}
