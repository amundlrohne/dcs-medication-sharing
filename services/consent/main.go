package main

import (
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// e.GET("/consent-provider/hello", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	configs.ConnectDB()
	routes.ConsentRoute(e)
	
	runCronJobs(e.AcquireContext())

	_ = e.Start(":8080")
}

func runCronJobs(c echo.Context) {
	s := cron.New()
    	s.AddFunc("@daily", func() {
		e :=echo.New()
		controllers.TimeOut(e.AcquireContext())
    })

	fmt.Println("Here")
    	s.Start()
   }
