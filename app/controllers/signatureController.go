package controllers

import (
	cb "casacomigo/app/controllers/base"
	//"casacomigo/app/services/account"
	"casacomigo/app/models/account"
	"github.com/robfig/revel"
    //"fmt"
)

type (
	Signature struct {
		cb.BaseController
	}
)

func init() {

}

//** CONTROLLER FUNCTIONS
func (s Signature) Register(noivo, telefoneNoivo, emailNoivo, noiva, telefoneNoiva, emailNoiva, apelido, senha, dataCasamento string) revel.Result {
	s.Validation.Required(noivo).Message("Nome do noivo.")
	s.Validation.Required(noiva).Message("Nome do noiva.")

	s.Validation.Required(telefoneNoivo).Message("Telefone do noivo.")
	s.Validation.Required(telefoneNoiva).Message("Telefone da noiva.")

	s.Validation.Required(emailNoivo).Message("Email do noivo.")
	s.Validation.Required(emailNoiva).Message("Email da noiva.")

	s.Validation.Required(apelido).Message("Email da noiva.")
	s.Validation.Required(dataCasamento).Message("Email da noiva.")

	s.Validation.MinSize(senha, 6).Message("Senha precisa ter 6 caracteres.")

	// Handle errors
	if s.Validation.HasErrors() {
    	s.Validation.Keep()
        s.FlashParams()
		return s.Redirect(s.Register)
	}

	a := account.Account{}
	a.Noivo 		= noivo
	a.Noiva 		= noiva
	a.TelefoneNoivo = telefoneNoivo
	a.TelefoneNoiva = telefoneNoiva
	a.EmailNoivo 	= emailNoivo
	a.EmailNoiva 	= emailNoiva
	a.Apelido 		= apelido
	a.Senha 		= senha
	a.Status 		= "aguardando_pagamento"
	a.Lucro 		= 0.0

	return s.Redirect(s.ChooseSite)
}

func (this *Signature) NewAccount() revel.Result {
	return this.Render("step_one.html");
}

func (this *Signature) ChooseSite() revel.Result {
	return this.Render();
}
