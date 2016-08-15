package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


func (this *MainController) Login() {
	this.TplName = "Account/signin.html"
}

func (this *MainController) Signin() {
	this.TplName = "Signature/newaccount.html"
}

func (this *MainController) Home() {
	this.TplName = "Home/index.html"
}
