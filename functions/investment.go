package main

// main describes the main entry point for your application

import (
	"fmt"
	"math"
)

func main() {
	//explicit type declaration
	// var investAmount, years float64 = 1000, 10
	//assignment with implicit type declaration
	// expectedReturnRate := 5.5

	// one liner assignment
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected yearly return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years Investing: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Real Value: ", futureRealValue)
	//fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureRealValue)

	fmt.Print(formattedFV)

}
