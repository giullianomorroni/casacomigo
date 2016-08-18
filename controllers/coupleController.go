package controllers

import (
	srv "casacomigo/services/account"
	"github.com/astaxie/beego"
)

type CoupleController struct {
	beego.Controller
}

func init() {}

func (this *CoupleController) FindCouple() {
	noivo := this.Input().Get("noivo")
	noiva := this.Input().Get("noiva")
	beego.Info(noivo)
	result := srv.FindCouple(noivo, noiva);
	this.Data["result"] = result
	this.TplName = "Couple/findcouple.html"
}
