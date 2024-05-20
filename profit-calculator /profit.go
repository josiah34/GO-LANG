package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Welcome to the Profit Calculator!")

	fmt.Print("Please enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter your total expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter your tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := revenue - expenses - (revenue * taxRate / 100)
	profitRatio := profit / revenue * 100
	fmt.Println("************************************")
	fmt.Println("Your earnings before taxes are: $", math.Round(ebt*100)/100)
	fmt.Println("Your profit is: $", math.Round(profit*100)/100)
	fmt.Println("Your profit ratio is: ", math.Round(profitRatio*100)/100, "%")

}
