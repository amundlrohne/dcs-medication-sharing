package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/consent/controllers"
	"github.com/labstack/echo/v4"
)

func ConsentRoute(e *echo.Echo) {
	e.POST("/consent", controllers.CreateConsent)
	e.GET("/consent/:from_public_key", controllers.GetConsent)
	e.GET("/consent", controllers.GetAllConsents)
	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")
}
