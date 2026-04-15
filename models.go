package main

import (
	"fmt"
	"time"
)

func (u User) print() {
	fmt.Println("Height: ", u.Height)
	fmt.Println("Weight: ", u.Weight)
}

func (u User) bmi(height, weight float64) {
	bmi := weight / (height * height)
	switch {
	case 0 < bmi && bmi < 18.5:
		fmt.Printf("Your BMI is %.2f! It's underweight! Go to doctor!\n", bmi)
	case 18.5 <= bmi && bmi <= 24.9:
		fmt.Printf("Your BMI is %.2f it's normal! Congrats!\n", bmi)
	case 25 <= bmi && bmi <= 29.9:
		fmt.Printf("Yout BMI is %.2f it's overweight! Careful!\n", bmi)
	case 30 <= bmi && bmi <= 34.9:
		fmt.Printf("Your BMI is %.2f it's obese! Full back!\n", bmi)
	case bmi >= 35:
		fmt.Printf("Your BMI is %.2f! It's extrimely obese! Go to doctor!\n", bmi)
	default:
		fmt.Printf("Error! You are eather dead or not born yet!")
	}
}

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
