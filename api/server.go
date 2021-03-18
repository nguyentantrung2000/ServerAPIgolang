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

type Students []Student

var students Students

func CreateStudent(c echo.Context) error {
	// Get name and email
	student := Student{}
	if err := c.Bind(&student); err != nil {
		fmt.Println(err)
		return err
	}
	students = append(students, student)
	fmt.Println(student)

	return nil
	// Id := c.FormValue("id")
	// Name := c.FormValue("Name")
	// GPA := c.FormValue("GPA")
	// IsMale := c.FormValue("IsMale")
	// return c.String(http.StatusOK, "id:"+Id+", name:"+Name+", gpa:"+GPA+", ismale:"+IsMale)
}

func GetStudentid(c echo.Context) error {
	id := c.QueryParam("id")
	for _, value := range students {
		if value.Id == id {
			return c.JSON(http.StatusOK, value)
		}
	}
	return c.JSON(http.StatusBadRequest, false)
}
func DeleteStudentid(c echo.Context) error {
	id := c.QueryParam("id")
	for i, _ := range students {
		if students[i].Id == id {
			students = append(students[:i], students[i+1:]...)
			return c.JSON(http.StatusOK, students)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func GetStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}
