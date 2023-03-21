package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/consent/controllers"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/responses"
	"github.com/labstack/echo/v4"
)

func ConsentRoute(e *echo.Echo) {
	e.POST("/createConsent", controllers.CreateConsent)
	e.GET("/getConsent/:from_public_key", controllers.GetConsent)
	e.GET("/getConsents", controllers.GetAllConsents)
	e.GET("/isalive", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.JSON(http.StatusOK, responses.ConsentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "OK"}})
}
