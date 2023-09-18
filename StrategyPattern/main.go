// Strategy Pattern includes
// 1. Context
// 2. Strategy
// 3. Concrete Strategies
package main

import (
	"fmt"
)

// Strategy
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete Strategies
type CreditCardPayment struct{}
type PayPalPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

func (pp *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
}

// Context
type ShoppingCart struct {
	PaymentMethod PaymentStrategy
}

func (cart *ShoppingCart) SetPaymentMethod(paymentMethod PaymentStrategy) {
	cart.PaymentMethod = paymentMethod
}

func (cart *ShoppingCart) Checkout(amount float64) {
	if cart.PaymentMethod == nil {
		fmt.Println("Please set a payment method first")
		return
	}
	cart.PaymentMethod.Pay(amount)
}

func main() {
	// Contructor Context
	cart := ShoppingCart{}

	// Choose Concrete Strategies
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}

	cart.SetPaymentMethod(creditCard)
	cart.Checkout(100.50)

	cart.SetPaymentMethod(payPal)
	cart.Checkout(55.25)
}
