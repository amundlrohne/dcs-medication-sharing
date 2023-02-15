package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/medication-record/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	_ = e.Start(":8080")
}
