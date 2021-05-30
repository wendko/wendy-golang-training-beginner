package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/controller"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/storage"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/payment-codes", controller.GetPaymentCodes)
	e.POST("/payment-codes", controller.CreatePaymentCode)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}