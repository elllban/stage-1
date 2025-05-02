package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"stage-1/internal/db"
	"stage-1/internal/handlers/taskHandlers"
	"stage-1/internal/handlers/userHandlers"
	"stage-1/internal/repository/taskRepository"
	"stage-1/internal/repository/userRepository"
	"stage-1/internal/service/taskService"
	"stage-1/internal/service/userService"
	"stage-1/internal/web/tasks"
	"stage-1/internal/web/users"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	tsRepo := taskRepository.NewTaskRepository(database)
	tsService := taskService.NewTaskService(tsRepo)
	tsHandlers := taskHandlers.NewTaskHandler(tsService)

	usRepo := userRepository.NewUserRepository(database)
	usService := userService.NewUserService(usRepo)
	usHandlers := userHandlers.NewUserHandler(usService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	tsStrictHandler := tasks.NewStrictHandler(tsHandlers, nil)
	tasks.RegisterHandlers(e, tsStrictHandler)

	usStrictHandler := users.NewStrictHandler(usHandlers, nil)
	users.RegisterHandlers(e, usStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
