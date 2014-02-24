package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/product"
	"github.com/robfig/revel/cache"
	"casacomigo/app/models/kart"
	"github.com/robfig/revel"
	//"fmt"
)

type (
	Product struct {
		cb.BaseController
	}
)

func init() {}

func (this *Product) List(limite, pular int) revel.Result {	
	list := srv.List(limite, pular);
	return this.RenderJson(list);
}

func (this *Product) Detail(codigo int) revel.Result {	
	product := srv.FindProduct(codigo);
	return this.RenderJson(product);
}

func (this *Product) Kart(codigo, quantidade int) revel.Result {	
	userKart := new(kart.Kart);
	err := cache.Get(this.Session.Id(), &userKart)
	if (err != nil) {
		userKart = new(kart.Kart);
	}
	userKart.AddItem(codigo, quantidade);
	cache.Set(this.Session.Id(), userKart, cache.DEFAULT)
	return this.Render();
}

func (this *Product) ShowKart() revel.Result {	
	userKart := new(kart.Kart);
	cache.Get(this.Session.Id(), &userKart)
	return this.RenderJson(userKart);
}

func (this *Product) ProductsByCategory(category string) revel.Result {	
	return this.Redirect(Product.List());
}