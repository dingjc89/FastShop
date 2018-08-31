package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 用户验证
	beego.Router("/", &controllers.LoginController{}, "get:Login")
	beego.Router("/loginin", &controllers.LoginController{}, "post:LoginIn")
	beego.Router("/logout", &controllers.LoginController{}, "get:LoginOut")

	// 主页
	beego.Router("/home", &controllers.HomeController{}, "get:Index")

	beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&controllers.RoleController{})

}
