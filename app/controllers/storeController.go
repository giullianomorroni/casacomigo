package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/product"
	mdl "casacomigo/app/models/product"
	"casacomigo/app/models/kart"
	"github.com/robfig/revel"
	"github.com/robfig/revel/cache"
	"time"
	//"fmt"
)

type (
	Store struct {
		cb.BaseController
	}
)

func init() {}


func (this *Store) Offers() revel.Result {
	return this.Render();
}

func (this *Store) Store(apelido string) revel.Result {
	this.Session["couple"] = apelido;
	return this.Render();
}

func (this *Store) Checkout() revel.Result {
	var allItens [20]kart.Kart;
	cache.Get(this.Session.Id(), &allItens)
	
	revel.TRACE.Printf("allItens %s", allItens)
	var products = []mdl.Product{}
	for x:=0; x < len(allItens); x++ {
		codigo := allItens[x].CodigoProduto;
		if codigo == 0 { continue; }
		product := srv.FindProduct(codigo);

		newProducts := make([]mdl.Product, 1)
		newProducts[0] = product;

		copy(newProducts[0:], products)
	}
	return this.Render(products);
}

func (this *Store) Faq() revel.Result {
	return this.Render();
}

func (this *Store) Contact() revel.Result {
	return this.Render();
}

func (this *Store) AddToKart(codigo, quantidade int) revel.Result {	
	item := kart.Kart{codigo, quantidade};
	var allItens [20]kart.Kart;
	err := cache.Get(this.Session.Id(), &allItens)
	if (err != nil) { }
	
	var cont = 0;
	for cont < 20 {
		if (allItens[cont].CodigoProduto == 0) {
			allItens[cont] = item;
			break;
		}
		cont++;
	}

	cache.Set(this.Session.Id(), allItens, 20*time.Minute)
	return this.Render();
}

func (this *Store) ShowKart() revel.Result {	
	var allItens [20]kart.Kart;
	cache.Get(this.Session.Id(), &allItens)
	return this.RenderJson(allItens);
}
