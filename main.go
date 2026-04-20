package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/skljor/day-checker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	var user models.User

	if err := db.First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		registerNewUser(db)
	} else if err != nil {
		//if something else goes wrong
		log.Fatalf("Database error: %v", err)
	} else {
		handleExistingUser(db, user)
	}
}

// initdb for declaring database
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.WeightEntry{})
	return db
}

// func for register new user
func registerNewUser(db *gorm.DB) {
	var newUser models.User
	fmt.Println("Welcome! Let's create your profile")

	for {
		fmt.Print("What is your height(m) and weight(kg): ")

		if _, err := fmt.Scan(&newUser.Height, &newUser.Weight); err != nil {
			fmt.Println("Parameters should be numerical")
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		if newUser.Height <= 0 || newUser.Weight <= 0 {
			fmt.Println("Error! Height and weight must be greater than zero.")
			continue
		}
		break
	}

	db.Create(&newUser)
	db.Create(&models.WeightEntry{Weight: newUser.Weight, UserID: newUser.ID})

	fmt.Println("Profile is successfully created!")

	bmi, status := newUser.BMI(newUser.Weight)
	fmt.Printf("Your height is %.2f m, weight is %.2f kg\n", newUser.Height, newUser.Weight)
	fmt.Printf("Your initial BMI is %.2f (%s)\n", bmi, status)
}

// func for exicting user
func handleExistingUser(db *gorm.DB, user models.User) {
	var lastWeight models.WeightEntry

	db.Where("user_id = ?", user.ID).Order("id desc").First(&lastWeight)

	displayWeight := user.Weight
	if lastWeight.ID != 0 {
		displayWeight = lastWeight.Weight
	}

	bmi, status := user.BMI(displayWeight)
	fmt.Printf("Welcome back! Your current weight is %.2f\n", displayWeight)
	fmt.Printf("Your index: %.2f. You are %s\n", bmi, status)

	var newWeight float64
	fmt.Print("Enter your current weight or press 0 to skip: ")
	fmt.Scanln(&newWeight)

	if newWeight > 0 {
		db.Create(&models.WeightEntry{
			Weight: newWeight,
			UserID: user.ID,
		})
		fmt.Println("Record saved")
	}
}
