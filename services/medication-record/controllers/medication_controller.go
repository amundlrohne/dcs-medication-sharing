package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//var medicationCollection *mongo.Collection = configs.GetCollection(configs.DB, "medication")
//var medicationCollection

func GetMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "GET medication-record/"+id+" Hey2")
}

func PostMedication(c echo.Context) error {
	return c.String(http.StatusOK, "POST medication-record/ testing...")
}

func PutMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "PUT medication-record/"+id)
}

func DeleteMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "DELETE medication-record/"+id)
}
