package controllers

import (
    "echo-mongo-api/configs"
    "echo-mongo-api/models"
    "echo-mongo-api/responses"
    "net/http"
    "time"
  
    "github.com/go-playground/validator/v10"
    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/net/context"
)

func getMedication(c echo.Context) error {
	id:=c.Param("id");
	return c.String(http.StatusOK, "GET medication-record/" + id);
}

func postMedication(c echo.Context) error {
	return c.String(http.StatusOK, "POST medication-record/");
}

func putMedication(c echo.Context) error {
	id:=c.Param("id");
	return c.String(http.StatusOK, "PUT medication-record/" + id);
}

func deleteMedication(c echo.Context) error {
	id:=c.Param("id");
	return c.String(http.StatusOK, "DELETE medication-record/" + id);
}
