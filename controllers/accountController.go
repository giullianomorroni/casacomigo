package controllers

import (
	srv "casacomigo/services/account"
	"casacomigo/models/account"
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}


func init() {}

func (this *AccountController) Login() {
	apelido := this.Input().Get("apelido")
	senha := this.Input().Get("senha")
	beego.Info(apelido)
	account := srv.Login(apelido, senha);
	this.SetSession("account", account.Apelido)
	this.Data["account"] = account
	this.Profile()
}

func (this *AccountController) Profile() {
	apelido := this.GetSession("account").(string)

	act := srv.FindAccount(apelido);
	ntfs := srv.FindNotifications(apelido);
	bnks := srv.FindBanks(apelido);

	profile := account.Profile{}
	profile.Account = act;
	profile.Notifications = ntfs;
	profile.Banks = bnks;
	this.Data["profile"] = profile
	this.TplName = "Account/profile.html"
}
