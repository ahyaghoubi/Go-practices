package main

import "fmt"

func main() {
	revenue := printAndAssign("Enter the revenue amount: ")
	expenses := printAndAssign("Enter the expenses amount: ")
	taxRate := printAndAssign("Enter the tax rate: ")

	earningBeforeTax, earningAfterTax, ratio := Calculator(revenue, expenses, taxRate)

	fmt.Printf("Your earning before tax is: %.2f\n", earningBeforeTax)
	fmt.Printf("Your earning after tax is: %.2f\n", earningAfterTax)
	fmt.Printf("Your earning ratio is: %.2f\n", ratio)
}

func printAndAssign(text string) (result float64) {
	fmt.Print(text)
	fmt.Scan(&result)

	return result
}

func Calculator(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt - ebt*taxRate/100
	ratio = ebt / eat

	return ebt, eat, ratio
}
