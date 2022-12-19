package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   int
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "okiebro", Age: 15},
	{ID: 2, Name: "nuxt", Age: 16},
	{ID: 3, Name: "sveltekit", Age: 17},
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func getAllUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

type Err struct {
	Message string `json:"message"`
}

func createUser(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	users = append(users, u)

	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// endpoint /health ไม่ได้ต้องการ authen แต่ get, create อันนี้ต้องการ ก็เป็นตัวที่ group ไป
	e.GET("/health", healthHandler)

	g := e.Group("/api")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "okiebro" && password == "welp4455" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/users", getAllUser)
	g.POST("/users", createUser)

	log.Println("server started at default port, :8080")
	log.Fatal(e.Start(":8080"))
	log.Println("close")
}
