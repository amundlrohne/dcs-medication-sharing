package routes

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/controllers"
	"github.com/labstack/echo/v4"
)

func MedicationRoute(e *echo.Echo) {
	e.GET("/standardization/valid/:drugName", controllers.DrugExists)
	e.GET("/standardization/:drugName", controllers.SearchDrug)
}
