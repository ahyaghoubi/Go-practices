package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investment float64
	var years float64
	var expectedReturnRate float64
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investment)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	futureValue := investment * math.Pow((1+expectedReturnRate/100), years)
	inflationAdjustedValue := futureValue / math.Pow((1+inflationRate/100), years)
	fmt.Printf("The future value of the investment is $%.2f\n", futureValue)
	fmt.Printf("The inflation-adjusted value of the investment is $%.2f\n", inflationAdjustedValue)
}
