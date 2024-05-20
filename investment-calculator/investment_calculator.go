package main

// import statement to include fmt and math
import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64
	var investmentAmount float64
	var years float64
	var expectedReturn float64
	// := both declares and initializes a variable.
	// = only assigns a value to an already declared variable.
	// Scope:

	// := can only be used inside functions.
	// = can be used inside and outside functions.
	// Multiple Variables:

	// := can declare and initialize multiple variables in one line.
	// = assigns values to already declared variables.

	//*****************************************************************************
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	// asking for expected return rate
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturn)
	// asking for number of years
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	// Asking for inflation rate
	fmt.Print("Enter the inflation rate: ")
	fmt.Scan(&inflationRate)

	// Calculating the investment return and inflation adjusted return then printing the result
	investmentReturn := investmentAmount * math.Pow(1+expectedReturn/100, years)
	inflationAdjustedReturn := investmentReturn / math.Pow(1+inflationRate/100, years)
	fmt.Println("Return after ", years, " years is: $", math.Round(investmentReturn*100)/100)
	fmt.Println(inflationAdjustedReturn)
}
