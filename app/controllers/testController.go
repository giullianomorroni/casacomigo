package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
)

//** TYPES
type (
	Test struct {
		cb.BaseController
	}
)

//** INIT FUNCTION

func init() {

}

func (this *Test) PaymentApproved(param string) revel.Result {
	return this.Render(param);
}

func (this *Test) PaymentCanceled(param string) revel.Result {
	return this.Render(param);
}

func (this *Test) PayPalAccountConfirm() revel.Result {
	return this.Render(param);
	//https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=<tokenValue>
}