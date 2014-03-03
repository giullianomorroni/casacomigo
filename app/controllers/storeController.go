package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/product"
	mdl "casacomigo/app/models/product"
	accountSrv "casacomigo/app/services/account"
	"casacomigo/app/models/shopper"
	"casacomigo/app/models/kart"
	"casacomigo/app/services/paypal"
	"github.com/robfig/revel"
	"github.com/robfig/revel/cache"
	"time"
	"strconv"
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

	var products = [20]mdl.Product{}
	var count = 0;
	var total float64;
	for x:=0; x < len(allItens); x++ {
		codigo := allItens[x].CodigoProduto;
		if codigo == 0 { continue; }
		product := srv.FindProduct(codigo);
		//revel.TRACE.Printf("product %s", product)
		if product.Codigo != 0 { 
			products[x] = product
			total += product.Preco
			count = x+1;
		}
	}
	this.Session["total"] = strconv.FormatFloat(total, 'g', 1, 64);
	//revel.TRACE.Printf("products %s", products)
	var result = make([]mdl.Product, count)
	result = products[:count]
	//revel.TRACE.Printf("result %s", result)
	
	checkout := mdl.Checkout{result, total};
	
	return this.Render(checkout);
}

func (this *Store) Faq() revel.Result {
	return this.Render();
}

func (this *Store) Register() revel.Result {
	return this.Render();
}

func (this *Store) RegisterShopper(nome, email, telefone string) revel.Result {
	var casal = this.Session["couple"]
	var aux = this.Session["total"]
	this.Session["payment_type"] = "gift"
	total, err := strconv.ParseFloat(aux, 64);
	if (err != nil){
		//TODO
	}
	shopper := shopper.Shopper{}
	shopper.Nome = nome;
	shopper.Telefone = telefone;
	shopper.Email = email; 
	shopper.Total = total;
	shopper.Casal = casal;
	accountSrv.RegisterShopper(shopper)

	token, ack := paypal.GenerateToken(total, "Presente Casamento " + casal)
	if (ack != "Success") {
		return this.NotFound("Não foi possível completar a operação", shopper);
	}
	return this.Redirect("https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token="+token);
}

func (this *Store) Contact() revel.Result {
	return this.Render();
}

func (this *Store) Payment() revel.Result {
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
	return this.RenderJson("ok");
}
