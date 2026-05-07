package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/skljor/day-checker/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	var user models.User

	if err := db.First(&user).Error; err != nil {
		log.Fatal("There is no users in database, need to create profile")
	}

	r := gin.Default()

	r.GET("api/tasks", func(c *gin.Context) {
		tasks, err := models.GetUserTasks(db, user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't load tasks"})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	r.POST("api/tasks", func(c *gin.Context) {
		var inputTask models.Task

		if err := c.ShouldBindJSON(&inputTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data type"})
			return
		}

		inputTask.UserID = user.ID
		inputTask.Done = false

		if err := models.CreateTask(db, &inputTask); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't save a task"})
			return
		}

		c.JSON(http.StatusOK, inputTask)
	})

	r.PATCH("/api/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		taskID, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong ID format"})
			return
		}

		var input struct {
			Done bool `json:"done"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Wrond data type"})
			return
		}

		if err := models.ToggleTaskStatus(db, uint(taskID), input.Done); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't update task status"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
	})

	r.DELETE("/api/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		taskID, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		if err := models.DeleteTask(db, uint(taskID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't delete task"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
	})

	fmt.Println("Server is started! Open http://localhost:8080/api/tasks")
	r.Run(":8080")
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
