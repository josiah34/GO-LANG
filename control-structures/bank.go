package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(accountBalance float64) {
	os.WriteFile(accountBalanceFile, []byte(fmt.Sprintf("%.2f", accountBalance)), 0644)

}


func readBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
    balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
	
}
func main() {

	var accountBalance = readBalanceFromFile()
	fmt.Println("Welcome to GO Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		//checkBalance := choice == 1

		switch choice {
		case 1:
			fmt.Printf("Your balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")
				continue

			}
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds")
				continue
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount")
				continue
			} else {
				accountBalance -= withdrawAmount
				writeBalanceToFile(accountBalance)
				fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
			}

		default:
			fmt.Println("Goodbye!")
			return

		}

		// if choice == 1 {
		// 	fmt.Printf("Your balance is: $%.2f\n", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("Enter the amount to deposit: ")
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid deposit amount")
		// 		continue

		// 	}
		// accountBalance += depositAmount
		// fmt.Printf("Your new balance is: $%.2f\n", accountBalance)

		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("Enter the amount to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("Insufficient funds")
		// 		continue
		// 	}

		// 	if withdrawAmount <= 0 {
		// 		fmt.Println("Invalid withdraw amount")
		// 		continue
		// 	} else {
		// 		accountBalance -= withdrawAmount
		// 		fmt.Printf("Your new balance is: $%.2f\n", accountBalance)
		// 	}

		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }
	}
}
