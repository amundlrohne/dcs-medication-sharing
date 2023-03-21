package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/controllers"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Fetch drugs and create drug name array
	controllers.Drugs = controllers.ReadDrugsFile()
	controllers.DrugNames = controllers.CreateNamesList(controllers.Drugs)

	routes.StandardizationRoute(e)

	_ = e.Start(":8080")
}
