package main

import "fmt"

func main() {
	age := 20

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println("Adult Years:", *agePointer)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
