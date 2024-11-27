package main

// main describes the main entry point for your application

import (
	"fmt"
	"math"
)

// adding variable declaration at top level scope of the go file so it canbe accessed by all functions
const inflationRate float64 = 2.5

func main() {
	//explicit type declaration
	// var investAmount, years float64 = 1000, 10
	//assignment with implicit type declaration
	// expectedReturnRate := 5.5

	// one liner assignment
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected yearly return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years Investing: ")
	fmt.Scan(&years)

	futureValue := calculatFutureValue(investmentAmount, years, expectedReturnRate)
	futureRealValue := calculateRealValue(futureValue, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Real Value: ", futureRealValue)
	//fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureRealValue)

	fmt.Print(formattedFV)

	// outputTxt("Hello, World!")

}

// outputTxt is a function which outputs the string passed to it
func outputTxt(text string) {
	fmt.Print(string(text))
}

// return types must be explicitly declared in functions

func calculatFutureValue(investmentAmount float64, years float64, expectedReturnRate float64) float64 {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	return futureValue
}

func calculateRealValue(futurevalue float64, years float64) float64 {
	return futurevalue / math.Pow(1+inflationRate/100, years)
}

// calculateRealValueAlt calculates the real value of an investment after a specified number of years,
// taking into account a constant inflation rate. It returns the inflation-adjusted future value.
func calculateRealValueAlt(futurevalue float64, years float64) (rv float64) {
	rv = futurevalue / math.Pow(1+inflationRate/100, years)
	return 
}