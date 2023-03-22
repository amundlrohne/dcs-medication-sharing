package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/controllers"
	"github.com/labstack/echo/v4"
)

func ProviderRoute(e *echo.Echo) {
	e.POST("/healthcare-provider", controllers.CreateProvider)
	e.GET("/healthcare-provider/:providerID", controllers.GetProvider)
	e.GET("/healthcare-provider/all", controllers.GetAllProviders)
	e.POST("/healthcare-provider/verify", controllers.VerifyUser)
	e.GET("/healthcare-provider/current", controllers.ReadAuthCookie)
	e.DELETE("/healthcare-provider", controllers.DeleteAuthCookie)

	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")
}
