package controllers

import (
	cb "casacomigo/app/controllers/base"
	srv "casacomigo/app/services/account"
	"casacomigo/app/services/paypal"
	"github.com/robfig/revel"
	"strconv"
)

type (
	Payment struct {
		cb.BaseController
	}
)

func init() {}

func (this *Payment) PaymentApproved(param string) revel.Result {
	return this.Render(param);
}

func (this *Payment) PaymentCanceled(param string) revel.Result {
	return this.Render(param);
}

func (this *Payment) PayPalAccountConfirm() revel.Result {
	return this.Render();
	//https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=<tokenValue>
}

func (this *Payment) PaymentReturn(token, PayerID string) revel.Result {
	ammount := this.Session["ammount"];
	ammountFloat, err := strconv.ParseFloat(ammount, 64);
	if (err != nil){
		//TODO
	}
	paypal.ConfirmPayment(token, PayerID, "USD", ammountFloat);

	paymentType := this.Session["payment_type"]

	if paymentType == "signature" {
		apelido := this.Session["account_id"]
		srv.UpdateAccountStatus(apelido, "pagamento_efetuado");
	} else if paymentType == "gift" {
		casal := this.Session["couple"]
		srv.UpdateCoupleAmmount(casal, ammountFloat);
	}

	
	return this.Render();
}
