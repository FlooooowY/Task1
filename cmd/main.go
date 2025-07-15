package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not conect to DB: %v", err)
	}
	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewTaskHandler(taskService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", taskHandlers.GetTasks)
	e.POST("/tasks", taskHandlers.PostTasks)
	e.PATCH("/tasks/:id", taskHandlers.PatchTasks)
	e.DELETE("/tasks/:id", taskHandlers.DeleteTasks)
	e.Start("localhost:8080")
}
