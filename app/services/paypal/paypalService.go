package paypal

import (
	"github.com/robfig/revel"
	"github.com/crowdmob/paypal"
)

/*
* generate token for a shop transaction
*/
func GenerateToken(amount float64, name string) (string,string) {
	revel.TRACE.Printf("Generate Token - amount: %s", amount)
	// An example to setup paypal express checkout for digital goods
	currencyCode := "BRL"
	isSandbox:= true
	returnURL:= "http://casacomigo.net/payment/return"
	cancelURL:= "http://casacomigo.net/payment/canceled"

	// sandbox credentials
	client := paypal.NewDefaultClient("giullianomorroni-facilitator_api1.gmail.com", "1390819253", "AGR75IF1giC-pWSpNNZemgHXSMWIA0Vl0c81i3stMYbQpiroX-k7fhaD", isSandbox)

	// Make a array of your digital-goods
	testGoods := []paypal.PayPalDigitalGood{paypal.PayPalDigitalGood{
		Name: name, 
		Amount: amount,
		Quantity: 1,
	}}

	// Sum amounts and get the token!
	response, err := client.SetExpressCheckoutDigitalGoods(
		paypal.SumPayPalDigitalGoodAmounts(&testGoods), 
		currencyCode, 
		returnURL,
		cancelURL,
		testGoods,
	)

	if err != nil {
	   panic(err)
	}

	mapResult := response.Values
	ack := mapResult["ACK"]
	token := mapResult["TOKEN"]

	return token[0], ack[0];
}

func ConfirmPayment(token, payerId, currency string, amount float64) {
	isSandbox := true
	revel.TRACE.Printf("ConfirmPayment token: %s", token);
	revel.TRACE.Printf("ConfirmPayment payerId: %s", payerId);
	revel.TRACE.Printf("ConfirmPayment amount: %s", amount )

	client := paypal.NewDefaultClient("giullianomorroni-facilitator_api1.gmail.com", "1390819253", "AGR75IF1giC-pWSpNNZemgHXSMWIA0Vl0c81i3stMYbQpiroX-k7fhaD", isSandbox)
	response, err := client.DoExpressCheckoutSale(token, payerId, currency, amount);	

	if err != nil {
		revel.TRACE.Printf("ConfirmPayment ERR : %s", err)
	}
	revel.TRACE.Printf("ConfirmPayment response: %s", response)
}
