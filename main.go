package main

import (
	"fmt"
)

func main() {
	var testUser User

	fmt.Println("Enter your height: ")
	fmt.Scanln(&testUser.Height)
	fmt.Println("Enter your weight: ")
	fmt.Scanln(&testUser.Weight)
	testUser.print()
	testUser.bmi(testUser.Height, testUser.Weight)
}
