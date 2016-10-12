package controllers

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
)

type FacebookController struct {
	beego.Controller
}

func (this *FacebookController) Get() {
	conf := &oauth2.Config{
		Scopes:       []string{"email"},
		RedirectURL:  "http://ec2-52-197-106-109.ap-northeast-1.compute.amazonaws.com/facebook",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	code := this.GetString("code")
	fmt.Println(code)
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		this.Data["ok"] = "認証エラー"
	} else {
		client := conf.Client(oauth2.NoContext, tok)
		result, _ := client.Get("https://graph.facebook.com/me")

		//json struct
		account := struct {
			Name string `json:"name"`
			Id string `json:"id"`
		}{}

		_ = json.NewDecoder(result.Body).Decode(&account)


		byteArray, _ := ioutil.ReadAll(result.Body)
		json := string(byteArray)
		fmt.Printf("%#v",json)
		this.Data["ok"] = "ようこそ " + account.Name 
	}

	this.TplName = "main.tpl"
}

