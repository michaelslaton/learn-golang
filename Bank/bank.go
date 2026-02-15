package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile () (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance")
	}

	return balance, nil
}

func writeBalanceToFile (balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main(){
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------------")
		// panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	
		// wantsCheckBalance := choice == 1

		switch choice {
			case 1:
				fmt.Println("Your balance is", accountBalance)

			case 2:
				fmt.Print("Your deposit? ")
				var depositAmount float64
				fmt.Scan(&depositAmount)
				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)

			case 3:
				fmt.Print("Amount to withdraw? ")
				var withdrawlAmount float64
				fmt.Scan(&withdrawlAmount)
				if withdrawlAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
				if withdrawlAmount > accountBalance {
					fmt.Println("Insufficient funds.")
					continue
				}
				accountBalance -= withdrawlAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)

			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank!")
				return
		}
	
	// 	if choice == 1 {
	// 		fmt.Println("Your balance is", accountBalance)
	// 	} else if choice == 2 {
	// 		fmt.Print("Your deposit? ")
	// 		var depositAmount float64
	// 		fmt.Scan(&depositAmount)
	// 		if depositAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			continue
	// 		}
	// 		accountBalance += depositAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 	} else if choice == 3 {
	// 		fmt.Print("Amount to withdraw? ")
	// 		var withdrawlAmount float64
	// 		fmt.Scan(&withdrawlAmount)
	// 		if withdrawlAmount <= 0 {
	// 			fmt.Println("Invalid amount. Must be greater than 0.")
	// 			continue
	// 		}
	// 		if withdrawlAmount > accountBalance {
	// 			fmt.Println("Insufficient funds.")
	// 			continue
	// 		}
	// 		accountBalance -= withdrawlAmount
	// 		fmt.Println("Balance updated! New amount:", accountBalance)
	// 	} else {
	// 		fmt.Println("Goodbye!")
	// 		// return
	// 		break
	// 	}
	}

	// fmt.Println("Thanks for choosing our bank!")
}