package routes

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/controllers"
	"github.com/labstack/echo/v4"
)

func MedicationRoute(e *echo.Echo) {

	//All routes related to medication comes here
	e.GET("/medication-record/:id", controllers.GetMedication)
	e.POST("/medication-record", controllers.PostMedication)
	e.PUT("/medication-record/:id", controllers.PutMedication)
	e.DELETE("/medication-record/:id", controllers.DeleteMedication)

}
