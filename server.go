package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/isa424/backend/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	// Add middlewares to the router
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect to the database
	query := fmt.Sprintf("%v:%v@/%v", handlers.User, handlers.Pass, handlers.DBName)
	db, err := sql.Open("mysql", query)

	// If connection is not established stop the application
	if err != nil {
		log.Fatal(err)
	}

	// Test pinging the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create the handler
	h := &handlers.Handler{DB: db}

	// Setup routes
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})
	e.GET("/users/:user_id", h.GetUser)
	e.GET("/users", h.GetUsers)
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:user_id", h.UpdateUser)
	e.DELETE("/users/:user_id", h.DeleteUser)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
