package controllers

import (
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/consent/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/consent/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
		DateCreated:   consent.DateCreated,
	}

	result, err := consentCollection.InsertOne(ctx, newConsent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.ConsentResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetConsent(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fromPublicKey := c.Param("from_public_key")
	var consent models.Consent
	defer cancel()

	// objId, _ := primitive.ObjectIDFromHex(fromPublicKey)

	err := consentCollection.FindOne(ctx, bson.M{"frompublickey": fromPublicKey}).Decode(&consent)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ConsentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": consent}})
}

func GetAllConsents(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var consents []models.Consent
	defer cancel()

	results, err := consentCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleConsent models.Consent
		if err = results.Decode(&singleConsent); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		consents = append(consents, singleConsent)
	}

	return c.JSON(http.StatusOK, responses.ConsentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": consents}})
}

func DeleteConsent(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	fromPublicKey := c.Param("from_public_key")
	// var consent models.Consent
	defer cancel()

	fmt.Println(fromPublicKey)
	result, err := consentCollection.DeleteOne(ctx, bson.M{"frompublickey": fromPublicKey})

    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    if result.DeletedCount < 1 {
        return c.JSON(http.StatusNotFound, responses.ConsentResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Consent not found!"}})
    }

    return c.JSON(http.StatusOK, responses.ConsentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Consent successfully deleted!"}})
}

func TimeOut(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	expdate := (time.Now().Format("02-01-2023"))

	defer cancel()

	fmt.Println("Time Now:")
	fmt.Println(expdate)
	result, err := consentCollection.DeleteMany(ctx, bson.M{"expdate": expdate})

    if err != nil {
        return c.JSON(http.StatusInternalServerError, responses.ConsentResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
    }

    if result.DeletedCount < 1 {
        return c.JSON(http.StatusNotFound, responses.ConsentResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "Consent not found!"}})
    }

    return c.JSON(http.StatusOK, responses.ConsentResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "Consent successfully deleted!"}})

}
