package main

import (
	"fmt"

	"github.com/skljor/day-checker/models"
)

func main() {
	var testUser models.User

	fmt.Println("Enter your height: ")
	fmt.Scanln(&testUser.Height)
	fmt.Println("Enter your weight: ")
	fmt.Scanln(&testUser.Weight)
	testUser.Print()
	bmi, status := testUser.BMI()
	fmt.Printf("Your index: %.2f. You are %s\n", bmi, status)
}
