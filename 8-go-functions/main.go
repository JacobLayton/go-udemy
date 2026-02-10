package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15, 40)
	// Note that in this case the 1 is not included in the numbers slice because it is defined as a separate paramater "startingValue"
	fmt.Println("Sum: ", sum)

	numbers := []int{1, 10, 15}
	anotherSum := sumup(1, numbers...)
	// Only needs the 1 at the beginning because sumup requries startingValue
	fmt.Println("Another sum: ", anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	fmt.Println("Starting value: ", startingValue)
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}