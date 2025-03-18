package main

import "fmt"

// Define the PaymentGateway interface
type PaymentGateway interface {
    ProcessPayment(amount float64) string
    RefundPayment(transactionID string) string
}

// Define a struct for CreditCard payment method
type CreditCard struct {
    CardNumber string
    ExpiryDate string
}

// Define a struct for PayPal payment method
type PayPal struct {
    Email string
}

// Implement the ProcessPayment method for CreditCard
func (cc CreditCard) ProcessPayment(amount float64) string {
    return fmt.Sprintf("Processed payment of $%.2f using Credit Card ending with %s", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// Implement the RefundPayment method for CreditCard
func (cc CreditCard) RefundPayment(transactionID string) string {
    return fmt.Sprintf("Refunded payment with transaction ID %s using Credit Card", transactionID)
}

// Implement the ProcessPayment method for PayPal
func (pp PayPal) ProcessPayment(amount float64) string {
    return fmt.Sprintf("Processed payment of $%.2f using PayPal account: %s", amount, pp.Email)
}

// Implement the RefundPayment method for PayPal
func (pp PayPal) RefundPayment(transactionID string) string {
    return fmt.Sprintf("Refunded payment with transaction ID %s using PayPal", transactionID)
}

func main() {
    // Create instances of CreditCard and PayPal
    cc := CreditCard{
        CardNumber: "1234567890123456",
        ExpiryDate: "12/25",
    }
    pp := PayPal{
        Email: "user@example.com",
    }

    // Declare a variable of type PaymentGateway
    var paymentGateway PaymentGateway

    // Use CreditCard payment gateway
    paymentGateway = cc
    fmt.Println(paymentGateway.ProcessPayment(100.0))  // Process payment using CreditCard
    fmt.Println(paymentGateway.RefundPayment("TX12345")) // Refund payment using CreditCard

    // Use PayPal payment gateway
    paymentGateway = pp
    fmt.Println(paymentGateway.ProcessPayment(200.0))  // Process payment using PayPal
    fmt.Println(paymentGateway.RefundPayment("TX67890")) // Refund payment using PayPal
}
