package routers

import (
	"casacomigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/casal/buscar", &controllers.CoupleController{}, "get:FindCouple")
	beego.Router("/autenticar", &controllers.AccountController{}, "post:Login")
	beego.Router("/conta/registrar/", &controllers.SignatureController{}, "post:Register")
	beego.Router("/conta/selecionar/site", &controllers.SignatureController{}, "post:RegisterSite")
	beego.Router("/conta/registrar/pagamento", &controllers.SignatureController{}, "post:RegisterPayment")

	beego.Router("/loja/casal/", &controllers.StoreController{}, "get:Store")
	beego.Router("/loja/ofertas/", &controllers.StoreController{}, "get:Store")
	beego.Router("/loja/", &controllers.StoreController{}, "get:Store")


	beego.Router("/account", &controllers.AccountController{})
	beego.Router("/store", &controllers.StoreController{})
	beego.Router("/product", &controllers.ProductController{})
	beego.Router("/signature", &controllers.SignatureController{})

	//open templates
	beego.Router("/", &controllers.MainController{}, "get:Home")
	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/conta/nova", &controllers.MainController{}, "get:Signin")
	beego.Router("/loja/register", &controllers.MainController{}, "get:StoreRegister")
	beego.Router("/loja/faq", &controllers.MainController{}, "get:StoreFaq")
	beego.Router("/loja/contact", &controllers.MainController{}, "get:StoreContact")
	beego.Router("/loja/payment", &controllers.MainController{}, "get:StorePayment")


	beego.SetStaticPath("/static", "static")
}
