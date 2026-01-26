package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter duration (years): ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFv := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRfv := fmt.Sprintf("Future value (inflation adjusted): %.1f\n", futureRealValue)

	// fmt.Printf("Future value: %.0f\nFuture value (inflation adjusted): %.0f\n", futureValue, futureRealValue)
	fmt.Print(formattedFv, formattedRfv)
}