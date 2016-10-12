package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	this.TplName = "home.tpl"
}

