package routes

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/controllers"
	"github.com/labstack/echo/v4"
)

func ConsentRoute(e *echo.Echo) {
	e.POST("/createConsent", controllers.CreateConsent)
	e.GET("/getConsent/:from_public_key", controllers.GetConsent)
	e.GET("/getConsents", controllers.GetAllConsents)
	e.DELETE("/deleteConsent/:from_public_key", controllers.DeleteConsent)
	e.DELETE("/deleteTimeOut", controllers.TimeOut)
}
