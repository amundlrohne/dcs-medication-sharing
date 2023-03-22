package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	configs.ConnectDB()
	routes.ProviderRoute(e)

	_ = e.Start(":8080")
}
