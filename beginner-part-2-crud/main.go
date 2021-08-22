package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/controller"
	"github.com/wendko/wendy-golang-training-beginner/beginner-part-2-crud/storage"
)

func loadEnvVariables() {
	err := godotenv.Load(".env")
  
	if err != nil {
	  	log.Fatalf("Error loading .env file")
	}
}

func main() {
	// Load env variables
	loadEnvVariables()

	// Echo instance
	e := echo.New()
	storage.NewDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/payment-codes", controller.GetPaymentCodes)
	e.GET("/payment-codes/:id", controller.GetPaymentCode)
	e.POST("/payment-codes", controller.CreatePaymentCode)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}