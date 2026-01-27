package main

import (
	"fmt"
)

func main() {

	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Funds")
		fmt.Println("3. Withdraw Funds")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Deposit successful. New balance is: ", accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("Withdrawal successful. New balance is: ", accountBalance)
			}
		} else {
			fmt.Println("Okay, goodbye then!")
			break
		}
	}
	fmt.Println("Thank you for banking with us!")
}