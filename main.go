package main

import (
	"DemoServer1/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/createstudent", api.CreateStudent)
	e.GET("/student/:id", api.GetStudent)
	e.Logger.Fatal(e.Start(":1323"))
}
