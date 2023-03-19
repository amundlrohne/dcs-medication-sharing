package routes

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/controllers"
	"github.com/labstack/echo/v4"
)

func ProviderRoute(e *echo.Echo) {
	e.POST("/createProvider", controllers.CreateProvider)
	e.GET("/getProvider/:providerID", controllers.GetProvider)
	e.GET("/getProviders", controllers.GetAllProviders)
	e.POST("/verifyUser", controllers.VerifyUser)
	e.GET("/getCurrentUser", controllers.ReadAuthCookie)
	e.DELETE("/deleteAuthCookie", controllers.DeleteAuthCookie)
}
