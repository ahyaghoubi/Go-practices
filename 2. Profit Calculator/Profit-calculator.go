package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := printAndAssign("Enter the revenue amount: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	expenses, err := printAndAssign("Enter the expenses amount: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	taxRate, err := printAndAssign("Enter the tax rate: ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	earningBeforeTax, earningAfterTax, ratio := Calculator(revenue, expenses, taxRate)

	writeTofile("earningBeforeTax", earningBeforeTax)
	writeTofile("earningAfterTax", earningAfterTax)
	writeTofile("ratio", ratio)

	fmt.Printf("Your earning before tax is: %.2f\n", earningBeforeTax)
	fmt.Printf("Your earning after tax is: %.2f\n", earningAfterTax)
	fmt.Printf("Your earning ratio is: %.2f\n", ratio)
}

func writeTofile(title string, number float64) {
	data := fmt.Sprint(number)
	os.WriteFile(title, []byte(data), 0644)
}

func printAndAssign(text string) (result float64, err error) {
	fmt.Print(text)
	fmt.Scan(&result)

	if result <= 0 {
		return 0, errors.New("invalid value")
	}

	return result, nil
}

func Calculator(revenue, expenses, taxRate float64) (ebt, eat, ratio float64) {
	ebt = revenue - expenses
	eat = ebt - ebt*taxRate/100
	ratio = ebt / eat

	return ebt, eat, ratio
}
