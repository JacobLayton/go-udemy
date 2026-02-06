package main

import "fmt"

type product struct {
	id int
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Cooking", "Traveling", "Gaming"}
	fmt.Println("Hobbies: ", hobbies)
	fmt.Println("First element of hobbies: ", hobbies[0])
	newestHobbies := hobbies[1:]
	fmt.Println("Newest hobbies: ", newestHobbies)
	oldestHobbies := hobbies[:2]
	fmt.Println("Oldest hobbies: ", oldestHobbies)
	oldestHobbiesTwo := []string{}
	oldestHobbiesTwo = append(oldestHobbiesTwo, hobbies[0], hobbies[1])
	fmt.Println("Oldest hobbies two: ", oldestHobbiesTwo)
	oldestHobbies = oldestHobbies[1:3]
	fmt.Println("Resliced oldest hobbies: ", oldestHobbies)

	courseGoals := []string{"Learn go", "Have fun"}
	fmt.Println("Course goals: ", courseGoals)
	courseGoals[1] = "Master Arrays"
	courseGoals = append(courseGoals, "Build an app")
	fmt.Println("Updated course goals: ", courseGoals)

	products := []product{
		{title: "A Book", price: 10.99, id: 1},
		{title: "A Carpet", price: 100.99, id: 2},
	}
	fmt.Println("Products: ", products)
	products = append(products, product{title: "A Lamp", price: 14.99, id: 3})
	fmt.Println("Updated products: ", products)
}

// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println("Second price in array: ", prices[1])
// 	prices[1] = 20.99
// 	fmt.Println("Updated prices array: ", prices)

// 	prices = append(prices, 5.99) // Add new price on the end
// 	fmt.Println("Prices array after append: ", prices)

// 	prices = prices[1:] // remove first price in the array
// 	fmt.Println("Prices array after slicing first element: ", prices)
// }

// func main() {
// 	// instantiate array
// 	var productNames [4]string = [4]string{"A book"}
// 	// instantiate array, another way
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	productNames[2] = "A Carpet"
// 	fmt.Println("Product array:", productNames)
// 	fmt.Println("First product:", productNames[0])
// 	fmt.Println("Prices array:", prices)

// 	featuredPrices := prices[1:] // slice of prices array from first index to end
// 	fmt.Println("Featured prices:", featuredPrices)
// 	featuredPrices[0] = 199.99
// 	highlightedPrices := featuredPrices[:1] // slice from start to first index (not including)
// 	fmt.Println("Highlighted prices:", highlightedPrices)
// 	fmt.Println("Original prices :", prices) // note that line 17 update the original array
// 	fmt.Println("Highlighted price length: " ,len(highlightedPrices),"\nCapacity: ", cap(highlightedPrices)) // length and capacity of the slice

// 	highlightedPrices = highlightedPrices[:3] // expand slice to include more elements up to capacity
// 	fmt.Println("Expanded highlighted prices:", highlightedPrices)
// 	fmt.Println("Updated Highlighted price length: " ,len(highlightedPrices),"\nCapacity: ", cap(highlightedPrices)) // length and capacity of the slice
// 	// Important to note that slices can be expaned to the end but not to the beginning of the original array
// }
