package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/controllers"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	// Fetch drugs
	controllers.Drugs = controllers.ReadDrugsFile()

	routes.MedicationRoute(e)

	_ = e.Start(":8080")
}
