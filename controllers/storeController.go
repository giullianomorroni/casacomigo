package controllers

import (
	srv "casacomigo/services/product"
	mdl "casacomigo/models/product"
	accountSrv "casacomigo/services/account"
	"casacomigo/models/shopper"
	"casacomigo/models/kart"
	"casacomigo/services/paypal"
	"github.com/astaxie/beego"
	"strconv"
)

type StoreController struct {
	beego.Controller
}

func init() {}

func (this *StoreController) Store() {
	apelido := this.Input().Get("apelido")
	beego.Info(apelido)
	this.SetSession("couple", apelido);
	this.TplName = "Store/store.html"
}

func (this *StoreController) Checkout() {
	var allItens [20]kart.Kart;
	//cache.Get(this.Session.Id(), &allItens)

	var products = [20]mdl.Product{}
	var count = 0;
	var total float64;
	for x:=0; x < len(allItens); x++ {
		codigo := allItens[x].CodigoProduto;
		if codigo == 0 { continue; }
		product := srv.FindProduct(codigo);
		////revel.TRACE.Printf("product %s", product)
		if product.Codigo != 0 {
			products[x] = product
			total += product.Preco
			count = x+1;
		}
	}
	//this.Session["total"] = strconv.FormatFloat(total, 'g', 1, 64);
	////revel.TRACE.Printf("products %s", products)
	var result = make([]mdl.Product, count)
	result = products[:count]
	////revel.TRACE.Printf("result %s", result)

	checkout := mdl.Checkout{result, total};
	this.Data["checkout"] = checkout
	this.Render();
}


func (this *StoreController) RegisterShopper(nome, email, telefone string) {
	var casal = ""//this.Session["couple"]
	var aux = ""//this.Session["total"]
	//this.Session["payment_type"] = "gift"
	//this.Session["shopper"] = nome
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
		//this.NotFound("Não foi possível completar a operação", shopper);
	}
	this.Ctx.Redirect(302, "https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token="+token);
}

func (this *StoreController) AddToKart(codigo, quantidade int) {
	item := kart.Kart{codigo, quantidade};
	var allItens [20]kart.Kart;
	//err := cache.Get(this.Session.Id(), &allItens)
	//if (err != nil) { }

	var cont = 0;
	for cont < 20 {
		if (allItens[cont].CodigoProduto == 0) {
			allItens[cont] = item;
			break;
		}
		cont++;
	}

	//cache.Set(this.Session.Id(), allItens, 20*time.Minute)
	this.ServeJSON();
}
