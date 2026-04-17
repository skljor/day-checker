package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
}

type WeightEntry struct {
	gorm.Model
	Weight float64 `json:"weight"`
	UserID uint    `json:"user_id"`
}

type Task struct {
	gorm.Model
	Title    string `json:"title"`
	Category string `json:"category"`
	Done     bool   `json:"done"`
}

func (u User) Print() {
	fmt.Println("Height: ", u.Height)
	fmt.Println("Weight: ", u.Weight)
}

func (u User) BMI(currentWeight float64) (float64, string) {
	bmi := currentWeight / (u.Height * u.Height)
	var category string

	switch {
	case 0 <= bmi && bmi < 18.5:
		category = "underweight"
	case 18.5 <= bmi && bmi < 25:
		category = "normal"
	case 25 <= bmi && bmi < 30:
		category = "overweight"
	case 30 <= bmi && bmi < 35:
		category = "obese"
	case bmi >= 35:
		category = "extremely obese"
	default:
		category = "either dead or not born yet"
	}

	return bmi, category
}
