package main

import (
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
