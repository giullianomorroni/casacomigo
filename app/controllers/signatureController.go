package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/account"
	mdl "casacomigo/app/models/account"
	"casacomigo/app/models/site"
	"casacomigo/app/services/paypal"
	"github.com/robfig/revel"
    "fmt"
)

type (
	Signature struct {
		cb.BaseController
	}
)

func init() {

}

//** CONTROLLER FUNCTIONS
func (s Signature) Register(nomeNoivo, telefoneNoivo, emailNoivo, nomeNoiva, telefoneNoiva, emailNoiva, apelido, senha, dataCasamento string) revel.Result {
	s.Validation.Required(nomeNoivo).Message("Nome do noivo.")
	s.Validation.Required(nomeNoiva).Message("Nome do noiva.")

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

	a := mdl.Account{}
	a.Noivo 		= nomeNoivo
	a.Noiva 		= nomeNoiva
	a.TelefoneNoivo = telefoneNoivo
	a.TelefoneNoiva = telefoneNoiva
	a.EmailNoivo 	= emailNoivo
	a.EmailNoiva 	= emailNoiva
	a.Apelido 		= apelido
	a.Senha 		= senha
	a.Status 		= "aguardando_pagamento"
	a.Lucro 		= 0.0

	p, err := srv.Register(a);

	fmt.Print(p)
	fmt.Print(err)
	
	return s.Redirect((*Signature).ChooseSite)
}

func (this *Signature) NewAccount() revel.Result {
	return this.Render("step_one.html");
}

func (this *Signature) ChooseSite() revel.Result {
	return this.Render();
}

func (this *Signature) Payment() revel.Result {
	return this.Render();
}

func (this *Signature) RegisterPayment(plano string) revel.Result {
	price := 0.0;
	if (plano == "trimestral") { price = 150.00; }
	if (plano == "semestral") { price = 250.00; }
	if (plano == "anual") { price = 350.00; }
	
	token, ack := paypal.GenerateToken(price, "Assinatura Casa Comigo")
	
	if (ack != "SUCESS") {
		//NOTHING
	}
	return this.Render(token);
}

func (this *Signature) RegisterSite(site_id string) revel.Result {
	siteChoose := site.Site{}
	siteChoose.Casal = "1"
	siteChoose.Site 	= site_id;

	p, err := srv.RegisterSite(siteChoose);

	fmt.Print(p)
	fmt.Print(err)

	return this.Redirect((*Signature).Payment)
}
