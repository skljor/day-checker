package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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
	for {
		fmt.Println("\n--- User menu ---")
		fmt.Println("1. Show current BMI and weight")
		fmt.Println("2. Update weight")
		fmt.Println("3. My tasks")
		fmt.Println("4. Add new task")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showStatus(db, user)
		case 2:
			updateWeight(db, user)
		case 3:
			handleListTask(db, user.ID)
		case 4:
			handleAddTask(db, user.ID)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Wrond choice! Try again!")
		}
	}
}

func handleAddTask(db *gorm.DB, userID uint) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What do you need to do today? ")
	scanner.Scan()
	title := scanner.Text()
	title = strings.TrimSpace(title)

	fmt.Print("Which category is it? ")
	scanner.Scan()
	category := scanner.Text()
	category = strings.TrimSpace(category)

	newTask := models.Task{
		Title:    title,
		Category: category,
		UserID:   userID,
		Done:     false,
	}

	if err := models.CreateTask(db, &newTask); err != nil {
		fmt.Printf("Error adding a new task: %v\n", err)
	} else {
		fmt.Println("Task successfully added!")
	}
}

func handleListTask(db *gorm.DB, userID uint) {
	tasks, err := models.GetUserTasks(db, userID)
	if err != nil {
		fmt.Printf("Error receiving tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("You have 0 tasks yet.")
		return
	}

	fmt.Println("\nYour current tasks:")
	for _, t := range tasks {
		status := " "
		if t.Done {
			status = "X"
		}
		fmt.Printf("[%s] ID: %d | %s (%s)\n", status, t.ID, t.Title, t.Category)
	}

	fmt.Print("\nEnter Task ID to toggle status (or 0 to go back): ")
	var id uint
	fmt.Scanln(&id)

	if id > 0 {
		for _, t := range tasks {
			if t.ID == id {
				err := models.ToggleTaskStatus(db, id, !t.Done)
				if err != nil {
					fmt.Printf("Error %v\n", err)
				} else {
					fmt.Println("Status updated!")
				}
				break
			}
		}
	}
}

func showStatus(db *gorm.DB, user models.User) {
	var lastWeight models.WeightEntry
	db.Where("user_id = ?", user.ID).Order("id desc").First(&lastWeight)

	displayWeight := user.Weight
	if lastWeight.ID != 0 {
		displayWeight = lastWeight.Weight
	}

	bmi, status := user.BMI(displayWeight)
	fmt.Printf("\nYour current weight is %.2f kg\n", displayWeight)
	fmt.Printf("Your BMI is %.2f (%s)\n", bmi, status)
}

func updateWeight(db *gorm.DB, user models.User) {
	var newWeight float64
	fmt.Print("Enter your new weight: ")
	if _, err := fmt.Scanln(&newWeight); err != nil || newWeight <= 0 {
		fmt.Println("Error: enter a correct number.")
		return
	}

	db.Create(&models.WeightEntry{
		Weight: newWeight,
		UserID: user.ID,
	})

	fmt.Println("Weight successfully updated!")
}
