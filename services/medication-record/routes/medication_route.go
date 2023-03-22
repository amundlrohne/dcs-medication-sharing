package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/controllers"
	"github.com/labstack/echo/v4"
)

func MedicationRoute(e *echo.Echo) {
	//All routes related to medication comes here
	e.POST("/medication-record", controllers.GetMedicationBundle)
	e.GET("/medication-record", controllers.GetAllMedicationBundles)
	e.POST("/medication-record/new", controllers.PostMedicaitonBundle)
	e.DELETE("/medication-record", controllers.DeleteMedicationBundle)

	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")

}
