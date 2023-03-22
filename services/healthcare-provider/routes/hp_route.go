package routes

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/controllers"
	"github.com/labstack/echo/v4"
)

func ProviderRoute(e *echo.Echo) {
	e.POST("/health-provider", controllers.CreateProvider)
	e.GET("/health-provider/:providerID", controllers.GetProvider)
	e.GET("/health-provider/all", controllers.GetAllProviders)
	e.POST("/health-provider/verify", controllers.VerifyUser)
	e.GET("/health-provider/current", controllers.ReadAuthCookie)
	e.DELETE("/health-provider", controllers.DeleteAuthCookie)

	e.GET("/", IsAlive)
}

func IsAlive(c echo.Context) error {

	return c.String(http.StatusOK, "Alive")
}
