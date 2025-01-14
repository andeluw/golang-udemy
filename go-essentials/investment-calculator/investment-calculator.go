package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64
	// years := 10.0
	// expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)	
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// formattedFV := fmt.Sprintf("Future value: %.0f", futureValue)

	formattedFV := fmt.Sprintf(`Future value: %.0f`, futureValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future value: %.0f\n", futureValue)
	fmt.Println(formattedFV)
	fmt.Println("Future Real Value: ", futureRealValue)
}

func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	// return fv, rfv
	return
}