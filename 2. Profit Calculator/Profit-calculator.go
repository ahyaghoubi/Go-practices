package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter the revenue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses amount: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	fmt.Printf("Your earning before tax is: %.2f\n", earningBeforeTax)

	earningAfterTax := earningBeforeTax - earningBeforeTax*taxRate/100
	fmt.Printf("Your earning after tax is: %.2f\n", earningAfterTax)

	ratio := earningBeforeTax / earningAfterTax
	fmt.Printf("Your earning ratio is: %.2f\n", ratio)
}
