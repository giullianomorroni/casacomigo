package paypal

import (
	"fmt"
	"github.com/crowdmob/paypal"
)

/*
* generate token for a shop transaction
*/
func GenerateToken(amount float64, name string) (string,string) {
	// An example to setup paypal express checkout for digital goods
	currencyCode := "USD"
	isSandbox:= true
	returnURL:= "http://casacomigo.net/test/payment/return"
	cancelURL:= "http://casacomigo.net/test/payment/canceled"

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
	   // handle error in charging
	}

	mapResult := response.Values
	ack := mapResult["ACK"]
	token := mapResult["TOKEN"]

	fmt.Print(mapResult)
	return token[0], ack[0];
}

func ConfirmPayment(token, payerId, currency string, amount float64) {
	isSandbox := true
	fmt.Print("token: " + token);
	fmt.Print("payerId: " + payerId);
	fmt.Print("ammount: %s", amount);
	client := paypal.NewDefaultClient("giullianomorroni-facilitator_api1.gmail.com", "1390819253", "AGR75IF1giC-pWSpNNZemgHXSMWIA0Vl0c81i3stMYbQpiroX-k7fhaD", isSandbox)
	response, err := client.DoExpressCheckoutSale(token, payerId, currency, amount);	
	fmt.Print(response);
	fmt.Print(err);
}
