package controllers

import (
	cb "casacomigo/app/controllers/base"
	"github.com/robfig/revel"
	"fmt"
)

type (
	Test struct {
		cb.BaseController
	}
)

func init() {}

func (this *Test) PaymentApproved(param string) revel.Result {
	return this.Render(param);
}

func (this *Test) PaymentCanceled(param string) revel.Result {
	return this.Render(param);
}

func (this *Test) PayPalAccountConfirm() revel.Result {
	return this.Render();
	//https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=<tokenValue>
}


func (this *Test) PaymentReturn(param string) revel.Result {
	fmt.Print(param)
	return this.Render(param);
}