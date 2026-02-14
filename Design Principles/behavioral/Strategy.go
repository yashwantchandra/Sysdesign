package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f with credit card", amount)
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f with PayPal", amount)
}

type Cart struct {
	strategy PaymentStrategy
}

func (c *Cart) SetPayment(s PaymentStrategy) {
	c.strategy = s
}

func (c *Cart) Checkout(amount float64) string {
	return c.strategy.Pay(amount)
}

func main() {
	cart := &Cart{}
	cart.SetPayment(CreditCard{})
	fmt.Println(cart.Checkout(99.99))

	cart.SetPayment(PayPal{})
	fmt.Println(cart.Checkout(49.99))
}
