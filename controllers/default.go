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

func (this *MainController) StoreFaq() {
	this.TplName = "Store/faq.html"
}

func (this *MainController) StoreRegister() {
	this.TplName = "Store/register.html"
}

func (this *MainController) StoreContact() {
	this.TplName = "Store/contact.html"
}

func (this *MainController) StorePayment() {
	this.TplName = "Store/payment.html"
}
