package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/account"
	"casacomigo/app/models/account"
	"github.com/robfig/revel"
)

type (
	Account struct {
		cb.BaseController
	}
)

func init() {}

func (this *Account) Signin() revel.Result {
	return this.Render();
}

func (this *Account) Login(apelido, senha string) revel.Result {
	account := srv.Login(apelido, senha);
	this.Session["account"] = account.Apelido;
	return this.Redirect("/minha_conta");
}

func (this *Account) Profile() revel.Result {
	apelido := this.Session["account"];
	
	act := srv.FindAccount(apelido);
	ntfs := srv.FindNotifications(apelido);
	bnks := srv.FindBanks(apelido);

	profile := account.Profile{}
	profile.Account = act;
	profile.Notifications = ntfs;
	profile.Banks = bnks;

	return this.Render(profile);
}

