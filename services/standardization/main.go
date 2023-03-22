package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/controllers"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/routes"

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

	// Fetch drugs and create drug name array
	controllers.Drugs = controllers.ReadDrugsFile()
	controllers.DrugNames = controllers.CreateNamesList(controllers.Drugs)

	routes.StandardizationRoute(e)

	_ = e.Start(":8080")
}
