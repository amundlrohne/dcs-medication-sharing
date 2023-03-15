package routes

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/controllers"
	"github.com/labstack/echo/v4"
)

func ConsentRoute(e *echo.Echo) {
	e.POST("/consent", controllers.CreateConsent)
}
