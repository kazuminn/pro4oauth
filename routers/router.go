package routers

import (
	"pro4oauth/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/callback", &controllers.CallbackController{})
	beego.Router("/facebook", &controllers.FacebookController{})
}

