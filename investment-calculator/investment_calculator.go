package main

// import statement to include fmt and math
import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years float64 = 1000 , 10
	var expectedReturn = 5.5
	// var years float64 = 10
	// When type should be inferred, use := instead of var
    example := "Hello, World!"

	var investmentReturn = investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Print("Return after ", years, " years is: $", math.Round(investmentReturn*100)/100, "\n")
	fmt.Println(example)
}
