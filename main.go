package main

import (
	"github.com/gin-gonic/gin"
)

type task struct {
	Name        string
	Description string
}

var tasks = []task{
	{
		Name:        "akaPMS",
		Description: "HLS product",
	},
	{
		Name:        "IaC",
		Description: "Internal DevOps project",
	},
}

func getAllTasks(c *gin.Context) {
	c.JSON(200, tasks)
}

func main() {

	server := gin.Default()

	server.GET("/tasks", getAllTasks)
	server.POST("/tasks")
	server.DELETE("/tasks/{}")

	server.Run("localhost:30080")
}
