package controllers

import (
	srv "casa_comigo/services/account"
	mdl "casa_comigo/models/account"
	"casa_comigo/models/site"
	"casa_comigo/services/paypal"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
    "fmt"
)

type SignatureController struct {
	beego.Controller
}

func (this *SignatureController) Register() {
	nomeNoivo := this.Input().Get("nomeNoivo")
	telefoneNoivo := this.Input().Get("telefoneNoivo")
	emailNoivo := this.Input().Get("emailNoivo")
	nomeNoiva := this.Input().Get("nomeNoiva")
	telefoneNoiva := this.Input().Get("telefoneNoiva")
	emailNoiva := this.Input().Get("emailNoiva")
	apelido := this.Input().Get("apelido")
	senha := this.Input().Get("senha")
	dataCasamento := this.Input().Get("dataCasamento")

	valid := validation.Validation{}

	valid.Required(nomeNoivo, "nomeNoivo").Message("Nome do noivo.")
	valid.Required(nomeNoiva, "nomeNoiva").Message("Nome do noiva.")

	valid.Required(telefoneNoivo, "telefoneNoivo").Message("Telefone do noivo.")
	valid.Required(telefoneNoiva, "telefoneNoiva").Message("Telefone da noiva.")

	valid.Required(emailNoivo, "emailNoivo").Message("Email do noivo.")
	valid.Required(emailNoiva, "emailNoiva").Message("Email da noiva.")

	valid.Required(apelido, "apelido").Message("Email da noiva.")
	valid.Required(dataCasamento, "dataCasamento").Message("Email da noiva.")

	valid.MinSize(senha, 6, "senha").Message("Senha precisa ter 6 caracteres.")

	// Handle errors
	flash := beego.NewFlash()
	if valid.HasErrors() {
        flash.Error("Existemm erros no seu cadastro")
		return
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

	this.SetSession("payment_type","signature")

	_, err := srv.Register(a);
	if (err != nil){
		flash.Error("Esse Apelido já está em uso, tente outro ;) !")
		return
	}

	this.SetSession("account_id", a.Apelido)
	this.TplName = "Signature/choosesite.html"
}

func (this *SignatureController) RegisterPayment() {
	plano := this.Input().Get("plano")
	price := 0.0;
	if (plano == "trimestral") { price = 150.00; }
	if (plano == "semestral") { price = 250.00; }
	if (plano == "anual") { price = 350.00; }

	//this.Session["ammount"] = strconv.FormatFloat(price, 'g', 1, 64);

	token, ack := paypal.GenerateToken(price, "Assinatura Casa Comigo")

	if (ack != "SUCESS") {
		//NOTHING
	}
	this.Data["token"] = token
	this.TplName = "Signature/registerpayment.html"	
}

func (this *SignatureController) RegisterSite() {
	site_id := this.Input().Get("site_id")

	siteChoose := site.Site{}
	siteChoose.Casal	= this.GetSession("account_id").(string)
	siteChoose.Site 	= site_id;

	p, err := srv.RegisterSite(siteChoose);

	fmt.Print(p)
	fmt.Print(err)
	this.TplName = "Signature/payment.html"
}
