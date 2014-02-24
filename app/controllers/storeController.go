package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
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

func (this *Store) Products() revel.Result {
	return this.Render();
}

func (this *Store) Checkout(product_id int) revel.Result {
	return this.Render();
}

func (this *Store) Product(codigo int) revel.Result {
	productCode := codigo;
	return this.Render(productCode);
}

func (this *Store) Faq() revel.Result {
	return this.Render();
}

func (this *Store) Contact() revel.Result {
	return this.Render();
}
