package controllers

import (
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/consent/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var consentCollection *mongo.Collection = configs.GetCollection(configs.DB, "consents")
var validate = validator.New()

func CreateConsent(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var consent models.Consent
	defer cancel()

	//validate the request body
	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ConsentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&consent); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.ConsentResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newConsent := models.Consent{
		ToPublicKey:   consent.ToPublicKey,
		FromPublicKey: consent.FromPublicKey,
		ExpDate:       consent.ExpDate,
	}

	result, err := consentCollection.InsertOne(ctx, newConsent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.ConsentResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}
