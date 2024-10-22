package main

import "fmt"

func main() {

}

type OnlineShop struct {
	payment PayPal
}

func NewOnlineShop(user string) *OnlineShop {
	return &OnlineShop{payment: PayPal{user: user}}
}

func (p *OnlineShop) Purchase(price float64) {
	p.payment.makePayment(price)
}

type PayPal struct {
	user string
}

func (p *PayPal) makePayment(price float64) {
	fmt.Println("PAYPAL => user:", p.user, "had a payment with price:", price)
}

type ZarinPal struct{}

func (p *ZarinPal) makePayment(user string, price float64) {
	fmt.Println("ZARIN-PAL => user:", user, "had a payment with price:", price)
}

// now how can we change our payment

// right struct
//type OnlineShop struct {
//	payment PaymentProcessor
//}
//
//func NewOnlineShop(processor PaymentProcessor) *OnlineShop {
//	return &OnlineShop{payment: processor}
//}
//
//func (p *OnlineShop) Purchase(price float64) {
//	p.payment.makePayment(price)
//}

type PaymentProcessor interface {
	makePayment(price float64)
}

type PayPalPaymentProcessor struct {
	paypal PayPal
}

func (p PayPalPaymentProcessor) makePayment(price float64) {
	p.paypal.makePayment(price)
}

type ZarinPalPaymentProcessor struct {
	zarinPal ZarinPal
	user     string
}

func (p ZarinPalPaymentProcessor) makePayment(price float64) {
	p.zarinPal.makePayment(p.user, price)
}
