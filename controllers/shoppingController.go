package controllers

import (
	"casa_comigo/services/paypal"
	"github.com/astaxie/beego"
)

type ShoppingController struct {
	beego.Controller
}

func init() {}

func (this *ShoppingController) ProductDetail(id int) {
	this.Render();
}

func (this *ShoppingController) Checkout(id int) {
	token, ack := paypal.GenerateToken(45.00, "Tv De Led Giulliano")
	if (ack != "SUCESS") {
		//NOTHING
	}
	this.Data["token"] = token
	this.Render();
}
