package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"stage-1/internal/db"
	"stage-1/internal/handlers"
	"stage-1/internal/repository"
	"stage-1/internal/service"
	"stage-1/internal/web/tasks"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	e := echo.New()

	tsRepo := repository.NewTaskRepository(database)
	tsService := service.NewTaskService(tsRepo)
	tsHandlers := handlers.NewTaskHandler(tsService)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictHandler := tasks.NewStrictHandler(tsHandlers, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
