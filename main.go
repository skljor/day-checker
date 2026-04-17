package main

import (
	"errors"
	"fmt"

	"github.com/skljor/day-checker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.WeightEntry{})

	var testUser models.User
	var lastWeight models.WeightEntry

	if err := db.First(&testUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Welcome! Let's create your profile")
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
		db.Create(&testUser)
		db.Create(&models.WeightEntry{Weight: testUser.Weight, UserID: testUser.ID})
		fmt.Println("Profile is succesfully created!")
		fmt.Printf("Your current height is %.2f m and weight is %.2f", testUser.Height, testUser.Weight)
	} else {
		db.Where("user_id = ?", testUser.ID).Order("id desc").First(&lastWeight)
		displayWeight := testUser.Weight
		if lastWeight.ID != 0 {
			displayWeight = lastWeight.Weight
		}

		bmi, status := testUser.BMI(displayWeight)
		fmt.Printf("Welcome back! Your current weight is %.2f kg\n", displayWeight)
		fmt.Printf("Your index: %.2f. You are %s\n", bmi, status)

		var newWeight float64
		fmt.Print("Enter your current weight or press 0 to skip: ")
		fmt.Scanln(&newWeight)

		if newWeight > 0 {
			db.Create(&models.WeightEntry{
				Weight: newWeight,
				UserID: testUser.ID,
			})
			fmt.Println("Record saved")
		}
	}

}
