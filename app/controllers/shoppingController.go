package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/account"
	"casacomigo/app/services/paypal"
	"github.com/robfig/revel"
	//"fmt"
)

type (
	Shopping struct {
		cb.BaseController
	}
)

func init() {}


func (this *Shopping) FindCouple(noivo, noiva string) revel.Result {
	result := srv.FindCouple(noivo, noiva);
	return this.Render(result);
}

func (this *Shopping) ProductDetail(id int) revel.Result {
	return this.Render();
}

func (this *Shopping) Checkout(id int) revel.Result {
	token, ack := paypal.GenerateToken(45.00, "Tv De Led Giulliano")
	if (ack != "SUCESS") {
		//NOTHING
	}
	return this.Render(token);
}