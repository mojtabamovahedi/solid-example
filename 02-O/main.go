package main

import "fmt"

func main() {

}

// bad way

type CreditCardPayment struct {
}

func (CreditCardPayment) ProcessPayment() {
	fmt.Println("paying with credit card")
}

type PayPalPayment struct {
}

func (PayPalPayment) ProcessPayment() {
	fmt.Println("paying with method pal")
}

type Order struct {
	way  string
	cPay CreditCardPayment
	pPay PayPalPayment
}

func (o Order) Pay() {
	if o.way == "credit_card_payment" {
		o.cPay.ProcessPayment()
	} else if o.way == "paypal_payment" {
		o.pPay.ProcessPayment()
	}
}

func NewOrderWithCreditPayment(c CreditCardPayment) *Order {
	return &Order{cPay: c, way: "credit_card_payment"}
}

func NewOrderWithPaypalPayment(p PayPalPayment) *Order {
	return &Order{pPay: p, way: "paypal_payment"}
}

// good way

//type PaymentMethod interface {
//	ProcessPayment()
//}
//
//type Order struct {
//	method PaymentMethod
//}
//func (o Order) Pay() {
//	o.method.ProcessPayment()
//}
