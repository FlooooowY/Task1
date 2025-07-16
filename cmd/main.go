package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"Tasks/internal/web/tasks"
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

	strictHandler := tasks.NewStrictHandler(taskHandlers, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
