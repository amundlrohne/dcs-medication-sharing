package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/controllers"
	"github.com/labstack/echo/v4"
)

func StandardizationRoute(e *echo.Echo) {
	e.GET("/standardization/drugNames/all", controllers.AllDrugNames)
	e.GET("/standardization/:drugName", controllers.SearchDrug)
	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")
}
