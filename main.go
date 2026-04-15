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
	bmi := testUser.Weight / (testUser.Height * testUser.Height)
	fmt.Printf("Your height: %v m, Your weight: %v kg, Your BMI %v \n", testUser.Height, testUser.Weight, bmi)

}
