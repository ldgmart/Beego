package routers

import (
	"github.com/astaxie/beego"
	controllers "sugoee/controllers"
)

func init() {
   beego.Router("/", &controllers.MainController{})
	beego.Router("/Signup/", &controllers.SignupController{})
	beego.Router("/Status/", &controllers.StatusController{}, "get:GetStatus")
	beego.Router("/Login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/task/", &controllers.TaskController{}, "get,post:NewSignup")
}
