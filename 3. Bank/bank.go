package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName string = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 1000, errors.New("unable to find the balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("unable to pasre to float64")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("---------")
	}

	var userChoice int
	fmt.Println("Welcome to GO bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Println("Your account balance:", accountBalance)
		} else if userChoice == 2 {
			var depositAmount float64
			fmt.Print("How much do you want to deposit? ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Print("Invalid amount. Please try again.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance:", accountBalance)

			writeBalanceToFile(accountBalance)
		} else if userChoice == 3 {
			var withdrawalAmount float64
			fmt.Print("How much do you want to withdraw? ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Print("Invalid amount. Please try again.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Your new balance:", accountBalance)

			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thankyou for choosing our bank!")
}
