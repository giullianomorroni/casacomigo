package routers

import (
	"casa_comigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/casal/buscar", &controllers.CoupleController{}, "get:FindCouple")
	beego.Router("/autenticar", &controllers.AccountController{}, "post:Login")
	beego.Router("/conta/registrar/", &controllers.SignatureController{}, "post:Register")
	beego.Router("/conta/selecionar/site", &controllers.SignatureController{}, "post:RegisterSite")
	beego.Router("/conta/registrar/pagamento", &controllers.SignatureController{}, "post:RegisterPayment")


	beego.Router("/account", &controllers.AccountController{})
	beego.Router("/store", &controllers.StoreController{})
	beego.Router("/product", &controllers.ProductController{})
	beego.Router("/signature", &controllers.SignatureController{})

	//open templates
	beego.Router("/", &controllers.MainController{}, "get:Home")
	beego.Router("/login", &controllers.MainController{}, "get:Login")
	beego.Router("/conta/nova", &controllers.MainController{}, "get:Signin")

	beego.SetStaticPath("/static", "static")
}
