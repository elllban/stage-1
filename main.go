package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func getTask(c echo.Context) error {
	response := "hello, " + task
	return c.JSON(http.StatusOK, response)
}

func postTask(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	task = req.Task
	return c.JSON(http.StatusCreated, task)
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)

	e.Start("localhost:8080")
}
