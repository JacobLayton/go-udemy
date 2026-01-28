package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------------")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		presentOptions()

		var choice int
		fmt.Print("your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be positive.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Deposit successful. New balance is: ", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
				continue
			} else {
				accountBalance -= withdrawAmount
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
				fmt.Println("Withdrawal successful. New balance is: ", accountBalance)
			}
		default:
			fmt.Println("Okay, goodbye then!")
			
			return
		}
	}
}