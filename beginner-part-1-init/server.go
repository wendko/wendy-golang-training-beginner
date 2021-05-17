
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

type HealthCheck struct {
	Status string `json:"status"`
}

func main() {
	e := echo.New()

	e.GET("/hello-world", func(c echo.Context) error {
		res := &Response{
			Message: "hello world",
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/health", func(c echo.Context) error {
		res := &HealthCheck{
			Status: "healthy",
		}
		return c.JSON(http.StatusOK, res)
	})


	e.Logger.Fatal(e.Start(":1323"))
}