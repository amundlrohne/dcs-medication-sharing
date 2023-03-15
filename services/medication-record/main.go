package main

import (
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/routes"
	"github.com/labstack/echo/v4"
)

func getMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "GET medication-record/"+id+"Heyyy Amund")
}

func postMedication(c echo.Context) error {
	return c.String(http.StatusOK, "POST medication-record/")
}

func putMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "PUT medication-record/"+id)
}

func deleteMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "DELETE medication-record/"+id)
}

func main() {

	e := echo.New()

	routes.MedicationRoute(e)

	e.Logger.Fatal(e.Start(":8080"))

}
