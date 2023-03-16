package controllers

import (
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/standardization/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var medicationCollection *mongo.Collection = configs.GetCollection(configs.DB, "medications")
var validate = validator.New()

func CreateMedication(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var medication models.Medication
	defer cancel()

	//validate the request body
	if err := c.Bind(&medication); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&medication); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newMedication := models.Medication{
		Id:    primitive.NewObjectID(),
		Name:  medication.Name,
		Dose:  medication.Dose,
		Title: medication.Title,
	}

	result, err := medicationCollection.InsertOne(ctx, newMedication)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.MedicationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.MedicationResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}
