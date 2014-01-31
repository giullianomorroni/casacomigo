package controllers

import (
	cb "casacomigo/app/controllers/base"
	"casacomigo/app/services/paypal"
	"github.com/robfig/revel"
)

type (
	Shopping struct {
		cb.BaseController
	}
)

func init() {}

func (this *Shopping) FindCouple() revel.Result {
	return this.Render();
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