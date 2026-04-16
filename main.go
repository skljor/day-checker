package main

import (
	"fmt"

	"github.com/skljor/day-checker/models"
)

func main() {
	var testUser models.User

	for {
		fmt.Print("What is your height(m) and weight(kg): ")

		if _, err := fmt.Scan(&testUser.Height, &testUser.Weight); err != nil {
			fmt.Println("Parameters should be numerical")
			//clear the input buffer
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		if testUser.Height <= 0 || testUser.Weight <= 0 {
			fmt.Println("Error: Height and weight must be greater than zero.")
			continue
		}

		break
	}
	testUser.Print()
	bmi, status := testUser.BMI()
	fmt.Printf("Your index: %.2f. You are %s\n", bmi, status)
}
