package models

import (
	"fmt"
	"time"
)

type User struct {
	Weight float64 `json:"weight"`
	Height float64 `json:"height"`
}

type WeightEntry struct {
	Date          time.Time `json:"date"`
	CurrentWeight float64   `json:"current_weight"`
}

type Task struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Done     bool   `json:"done"`
}

func (u User) Print() {
	fmt.Println("Height: ", u.Height)
	fmt.Println("Weight: ", u.Weight)
}

func (u User) BMI() (float64, string) {
	bmi := u.Weight / (u.Height * u.Height)
	var category string

	switch {
	case 0 < bmi && bmi < 18.5:
		category = "underweight"
	case 18.5 <= bmi && bmi <= 24.9:
		category = "normal"
	case 25 <= bmi && bmi <= 29.9:
		category = "overweight"
	case 30 <= bmi && bmi <= 34.9:
		category = "obese"
	case bmi >= 35:
		category = "extremely obese"
	default:
		category = "eather dead or not born yet"
	}

	return bmi, category
}
