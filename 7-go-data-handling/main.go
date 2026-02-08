package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5) // length 2, capacity 5

	// userNames = append(userNames, "Max")
	// userNames = append(userNames, "Manuel")
	// When using make to create a slice, appending will add after the empty slots, so we assign directly to the index
	userNames[0] = "Max"
	userNames[1] = "Manuel"
	userNames = append(userNames, "Jacob")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.6

	fmt.Println(courseRatings)
	// The point of utilizing the make function is to set the initial memory allocation for the slice or map, which
	// can improve performance by reducing the need for multiple memory allocations as elements are added. In this
	// example, we set the initial capacity of the userNames slice to 5, which means it can hold up to 5 elements
	// before needing to allocate more memory. Similarly, we set the initial size of the courseRatings map to 3,
	// which means it can hold up to 3 key-value pairs before needing to allocate more memory.
}
