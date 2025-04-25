package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var tasks []requestBody

type requestBody struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, &tasks)
}

func postTask(c echo.Context) error {
	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	ts := requestBody{
		ID:     uuid.NewString(),
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	tasks = append(tasks, ts)
	return c.JSON(http.StatusCreated, ts)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")

	var req requestBody
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	for i, ts := range tasks {
		if ts.ID == id {
			tasks[i].Task = req.Task
			tasks[i].IsDone = req.IsDone
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	for i, ts := range tasks {
		if ts.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Task not found"})
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)

	e.Start("localhost:8080")
}
