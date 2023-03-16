package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	routes.MedicationRoute(e)

	e.Logger.Fatal(e.Start(":8080"))

}
