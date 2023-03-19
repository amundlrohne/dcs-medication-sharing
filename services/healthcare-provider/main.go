package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDB()
	routes.ProviderRoute(e)

	_ = e.Start(":8080")
}
