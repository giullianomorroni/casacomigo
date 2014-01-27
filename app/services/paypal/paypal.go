package main

import (
  "fmt"
  "github.com/crowdmob/paypal"
)
   
func main() {
  paypalExpressCheckoutHandler()
}


func paypalExpressCheckoutHandler(w http.ResponseWriter, r *http.Request) {
  // An example to setup paypal express checkout for digital goods
  currencyCode := "USD"
  isSandbox    := true
  returnURL    := "http://casacomigo.com/test/payment/return"
  cancelURL    := "http://casacomigo.com/test/payment/canceled"

  // Create the paypal Client with default http client
  client := paypal.NewDefaultClient("giullianomorroni-facilitator_api1.gmail.com", "kaza6969", "AGR75IF1giC-pWSpNNZemgHXSMWIA0Vl0c81i3stMYbQpiroX-k7fhaD", isSandbox)

  // Make a array of your digital-goods
  testGoods := []paypal.PayPalDigitalGood{paypal.PayPalDigitalGood{
    Name: "Test Good", 
    Amount: 200.000,
    Quantity: 5,
  }}

  // Sum amounts and get the token!
  response, err := client.SetExpressCheckoutDigitalGoods(paypal.SumPayPalDigitalGoodAmounts(&testGoods), 
    currencyCode, 
    returnURL, 
    cancelURL, 
    testGoods,
  )

  if err != nil {
    // ... gracefully handle error
  } else { // redirect to paypal
    http.Redirect(w, r, response.CheckoutUrl(), 301)
  }
}