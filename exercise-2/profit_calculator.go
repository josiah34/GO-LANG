package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Please enter your revenue: ")

	expenses = getUserInput("Please enter your total expenses: ")

	taxRate = getUserInput("Please enter your tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(message string) float64 {
	var userInput float64
	fmt.Print(message)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

// func calculateEBT(revenue float64, expenses float64) float64 {
// 	return revenue - expenses
// }

// func calculateProfit(ebt float64, taxRate float64) float64 {
// 	return ebt * (1 - taxRate/100)
// }

// func calculateRatio(ebt float64, profit float64) float64 {
// 	return ebt / profit
// }
