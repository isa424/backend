package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func handler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct{ message string }{message: "OK"})
}

func main() {
	fmt.Println("Hello world")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":8080"))
}
