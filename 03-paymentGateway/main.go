package main

import "fmt"

type Payable interface {
	Pay(amount float64) error
}

type CreditCard struct {
	Number int
}

func (p CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with card %d.\n", amount, p.Number)
	return nil
}

type PayPal struct {
	Email string
}

func (c PayPal) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with Paypal %s.\n", amount, c.Email)
	return nil
}

func ProcessPayment(p Payable, amount float64) {
	p.Pay(amount)
}

func main() {
	myCreditCard := CreditCard{Number: 1122334455}
	myPayPal := PayPal{Email: "email@mail.com"}

	ProcessPayment(myCreditCard, 19.99)
	ProcessPayment(myPayPal, 35.86)
}
