package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"stage-1/internal/db"
	"stage-1/internal/handlers"
	"stage-1/internal/repository"
	"stage-1/internal/service"
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

	e.GET("/task", tsHandlers.GetTask)
	e.POST("/task", tsHandlers.PostTask)
	e.PATCH("/task/:id", tsHandlers.PatchTask)
	e.DELETE("/task/:id", tsHandlers.DeleteTask)

	e.Start("localhost:8080")
}
