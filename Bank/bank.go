package main

import "fmt"

func main(){
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
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