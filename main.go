package main

import (
	"DemoServer1/api"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/createstudent", api.CreateStudent)
	e.GET("/student", api.GetStudent)
	e.GET("/studentone", api.GetStudentid)
	e.DELETE("/deletestudent", api.DeleteStudentid)
	e.Logger.Fatal(e.Start(":1323"))
}
