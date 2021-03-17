package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Student struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	GPA    float32 `json:"gpa"`
	IsMale bool    `json:"ismale"`
}

func CreateStudent(c echo.Context) error {
	// Get name and email
	m := new(Student)
	if err := c.Bind(m); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(*m)

	return nil
	// Id := c.FormValue("id")
	// Name := c.FormValue("Name")
	// GPA := c.FormValue("GPA")
	// IsMale := c.FormValue("IsMale")
	// return c.String(http.StatusOK, "id:"+Id+", name:"+Name+", gpa:"+GPA+", ismale:"+IsMale)
}

func GetStudent(c echo.Context) error {
	Id := c.QueryParam("id")
	return c.String(http.StatusOK, "id:"+Id)
}
