package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// GET /tasks
func tasksGET(tasks *[]Task) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	}
}

// GET /tasks/:id
func taskGET(tasks *[]Task) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		for _, task := range *tasks {
			if task.ID == id {
				c.JSON(http.StatusOK, task)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

// POST /tasks
func createTask(tasks *[]Task, count *int) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTask Task
		if err := c.BindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Data"})
			return
		}
		newTask.ID = *count
		*count++
		*tasks = append(*tasks, newTask)
		c.JSON(http.StatusCreated, newTask)
	}
}

// PUT /tasks/:id
func updateTask(tasks *[]Task) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var updated Task
		if err := c.BindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		for i, task := range *tasks {
			if task.ID == id {
				(*tasks)[i].Title = updated.Title
				(*tasks)[i].Done = updated.Done
				c.JSON(http.StatusOK, (*tasks)[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
}

func main() {
	tasks := []Task{}
	taskCount := 1

	r := gin.Default()

	r.GET("/tasks", tasksGET(&tasks))
	r.GET("/tasks/:id", taskGET(&tasks))
	r.POST("/tasks", createTask(&tasks, &taskCount))
	r.PUT("/tasks/:id", updateTask(&tasks))

	r.Run(":8080")
}
