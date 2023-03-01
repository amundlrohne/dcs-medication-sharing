package routes

import "github.com/labstack/echo/v4"

func MedicationRoute(e *echo.Echo)  {
    //All routes related to medication comes here

    e.GET("/medication-record/:id", controllers.getMedication)
	e.POST("/medication-record/", controllers.postMedication)
	e.PUT("/medication-record/:id", controllers.putMedication)
	e.DELETE("/medication-record/:id", controllers.deleteMedication)

}