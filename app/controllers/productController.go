package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/product"
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

func (this *Product) Products() revel.Result {
	return this.Render();
}

func (this *Product) ProductsByCategory(category string) revel.Result {	
	return this.Redirect(this.List(10, 0));
}

func (this *Product) Product(codigo int) revel.Result {
	productCode := codigo;
	return this.Render(productCode);
}
