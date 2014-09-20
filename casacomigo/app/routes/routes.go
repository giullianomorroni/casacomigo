// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tBaseController struct {}
var BaseController tBaseController


func (_ tBaseController) Before(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.Before", args).Url
}

func (_ tBaseController) After(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.After", args).Url
}

func (_ tBaseController) Panic(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("BaseController.Panic", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tProduct struct {}
var Product tProduct


func (_ tProduct) List(
		limite int,
		pular int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "limite", limite)
	revel.Unbind(args, "pular", pular)
	return revel.MainRouter.Reverse("Product.List", args).Url
}

func (_ tProduct) ListOffers(
		limite int,
		pular int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "limite", limite)
	revel.Unbind(args, "pular", pular)
	return revel.MainRouter.Reverse("Product.ListOffers", args).Url
}

func (_ tProduct) Detail(
		codigo int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "codigo", codigo)
	return revel.MainRouter.Reverse("Product.Detail", args).Url
}

func (_ tProduct) Products(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Products", args).Url
}

func (_ tProduct) ProductsByCategory(
		category string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "category", category)
	return revel.MainRouter.Reverse("Product.ProductsByCategory", args).Url
}

func (_ tProduct) Product(
		codigo int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "codigo", codigo)
	return revel.MainRouter.Reverse("Product.Product", args).Url
}

func (_ tProduct) Offers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Product.Offers", args).Url
}


type tStore struct {}
var Store tStore


func (_ tStore) Store(
		apelido string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "apelido", apelido)
	return revel.MainRouter.Reverse("Store.Store", args).Url
}

func (_ tStore) Checkout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Store.Checkout", args).Url
}

func (_ tStore) Faq(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Store.Faq", args).Url
}

func (_ tStore) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Store.Register", args).Url
}

func (_ tStore) RegisterShopper(
		nome string,
		email string,
		telefone string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nome", nome)
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "telefone", telefone)
	return revel.MainRouter.Reverse("Store.RegisterShopper", args).Url
}

func (_ tStore) Contact(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Store.Contact", args).Url
}

func (_ tStore) Payment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Store.Payment", args).Url
}

func (_ tStore) AddToKart(
		codigo int,
		quantidade int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "codigo", codigo)
	revel.Unbind(args, "quantidade", quantidade)
	return revel.MainRouter.Reverse("Store.AddToKart", args).Url
}


type tAccount struct {}
var Account tAccount


func (_ tAccount) Signin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Account.Signin", args).Url
}

func (_ tAccount) Login(
		apelido string,
		senha string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "apelido", apelido)
	revel.Unbind(args, "senha", senha)
	return revel.MainRouter.Reverse("Account.Login", args).Url
}

func (_ tAccount) Profile(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Account.Profile", args).Url
}


type tPayment struct {}
var Payment tPayment


func (_ tPayment) PaymentApproved(
		param string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "param", param)
	return revel.MainRouter.Reverse("Payment.PaymentApproved", args).Url
}

func (_ tPayment) PaymentCanceled(
		param string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "param", param)
	return revel.MainRouter.Reverse("Payment.PaymentCanceled", args).Url
}

func (_ tPayment) PayPalAccountConfirm(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Payment.PayPalAccountConfirm", args).Url
}

func (_ tPayment) PaymentReturn(
		token string,
		PayerID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "token", token)
	revel.Unbind(args, "PayerID", PayerID)
	return revel.MainRouter.Reverse("Payment.PaymentReturn", args).Url
}


type tShopping struct {}
var Shopping tShopping


func (_ tShopping) FindCouple(
		noivo string,
		noiva string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "noivo", noivo)
	revel.Unbind(args, "noiva", noiva)
	return revel.MainRouter.Reverse("Shopping.FindCouple", args).Url
}

func (_ tShopping) ProductDetail(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Shopping.ProductDetail", args).Url
}

func (_ tShopping) Checkout(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Shopping.Checkout", args).Url
}


type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).Url
}


type tSignature struct {}
var Signature tSignature


func (_ tSignature) Register(
		nomeNoivo string,
		telefoneNoivo string,
		emailNoivo string,
		nomeNoiva string,
		telefoneNoiva string,
		emailNoiva string,
		apelido string,
		senha string,
		dataCasamento string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nomeNoivo", nomeNoivo)
	revel.Unbind(args, "telefoneNoivo", telefoneNoivo)
	revel.Unbind(args, "emailNoivo", emailNoivo)
	revel.Unbind(args, "nomeNoiva", nomeNoiva)
	revel.Unbind(args, "telefoneNoiva", telefoneNoiva)
	revel.Unbind(args, "emailNoiva", emailNoiva)
	revel.Unbind(args, "apelido", apelido)
	revel.Unbind(args, "senha", senha)
	revel.Unbind(args, "dataCasamento", dataCasamento)
	return revel.MainRouter.Reverse("Signature.Register", args).Url
}

func (_ tSignature) NewAccount(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Signature.NewAccount", args).Url
}

func (_ tSignature) ChooseSite(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Signature.ChooseSite", args).Url
}

func (_ tSignature) Payment(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Signature.Payment", args).Url
}

func (_ tSignature) RegisterPayment(
		plano string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "plano", plano)
	return revel.MainRouter.Reverse("Signature.RegisterPayment", args).Url
}

func (_ tSignature) RegisterSite(
		site_id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "site_id", site_id)
	return revel.MainRouter.Reverse("Signature.RegisterSite", args).Url
}


