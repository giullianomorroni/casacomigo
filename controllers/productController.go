package controllers

import (
	srv "casacomigo/services/product"
	"github.com/astaxie/beego"
	//"fmt"
)

type ProductController struct {
	beego.Controller
}

func init() {}

func (this *ProductController) List(limite, pular int) {
	list := srv.List(limite, pular);
	this.Data["list"] = list
	this.ServeJSON();
}

func (this *ProductController) ListOffers(limite, pular int) {
	list := srv.ListOffers(limite, pular);
	this.Data["list"] = list
	this.ServeJSON();
}

func (this *ProductController) Detail(codigo int) {
	product := srv.FindProduct(codigo);
	this.Data["product"] = product
	this.ServeJSON();
}

func (this *ProductController) Products() {
	this.Render();
}

func (this *ProductController) ProductsByCategory(category string) {
	this.Ctx.Redirect(200 , ""); //this.List(10, 0)
}

func (this *ProductController) Product(codigo int) {
	productCode := codigo;
	this.Data["productCode"] = productCode
	this.Render();
}

func (this *ProductController) Offers() {
	this.Render();
}
