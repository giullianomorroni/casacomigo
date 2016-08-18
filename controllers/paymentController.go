package controllers

import (
	srv "casacomigo/services/account"
	"casacomigo/services/paypal"
	"github.com/astaxie/beego"
	"strconv"
)

type PaymentController struct {
	beego.Controller
}

func init() {}

func (this *PaymentController) PaymentApproved(param string) {
	this.Data["param"] = param
	this.Render();
}

func (this *PaymentController) PaymentCanceled(param string) {
	this.Data["param"] = param
	this.Render();
}

func (this *PaymentController) PayPalAccountConfirm() {
	this.Render();
	//https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=<tokenValue>
}

func (this *PaymentController) PaymentReturn(token, PayerID string) {
	ammount := "200"//this.Session["ammount"];
	ammountFloat, err := strconv.ParseFloat(ammount, 64);
	if (err != nil){
		panic(err);
	}
	paypal.ConfirmPayment(token, PayerID, "USD", ammountFloat);

	paymentType := ""//this.Session["payment_type"]
	if paymentType == "signature" {
		apelido := ""//this.Session["account_id"]
		srv.UpdateAccountStatus(apelido, "pagamento_efetuado");
	} else if paymentType == "gift" {
		casal := ""//this.Session["couple"]
		shopper := ""//this.Session["shopper"]
		srv.UpdateCoupleAmmount(casal, shopper, ammountFloat);
	}

	this.Render();
}
