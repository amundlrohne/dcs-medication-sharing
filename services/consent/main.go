package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	configs.ConnectDB()
	routes.ConsentRoute(e)

	_ = e.Start(":8080")
}
