package main

import "fmt"

func main() {
	age := 32 // regular variable

	var userAge *int // asterisk indicates pointer type
	userAge = &age // pointer to age variable

	fmt.Println("Age:", age)

	fmt.Println("userAge pointer address:", userAge)
	fmt.Println("userAge dereferenced pointer value:", *userAge)

	adultYears := getAdultYears(userAge)
	fmt.Println("Adult Years:", adultYears)

	updateAgeToAdultYears(userAge)
	fmt.Println("After direct pointer update:", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func updateAgeToAdultYears(age *int) {
	*age = *age - 18
}