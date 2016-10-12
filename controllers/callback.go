package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
)

type CallbackController struct {
	beego.Controller
}

func (this *CallbackController) Get() {
	conf := &oauth2.Config{
		Scopes:       []string{"email"},
		RedirectURL:  "http://ec2-52-197-106-109.ap-northeast-1.compute.amazonaws.com/facebook",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)
	this.Redirect(url,302)
}

