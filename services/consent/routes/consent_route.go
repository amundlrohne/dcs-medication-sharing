package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/consent/controllers"
	"github.com/labstack/echo/v4"
)

func ConsentRoute(e *echo.Echo) {
	e.POST("/consent/", controllers.CreateConsent)
	e.GET("/consent/from/:from_public_key", controllers.GetConsentRequests)
	e.GET("/consent/to/:to_public_key", controllers.GetConsentIncoming)
	e.GET("/consent/", controllers.GetAllConsents)
	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")
}
