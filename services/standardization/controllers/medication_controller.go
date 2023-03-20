package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

// Global variables
// Struct for decoding drug json file
type Drug struct {
	Drug_name                     string
	Medical_condition             string
	Side_effects                  string
	Generic_name                  string
	Drug_classes                  string
	Brand_names                   string
	activity                      string
	Rx_otc                        string
	Pregnancy_category            string
	Csa                           string
	Alcohol                       string
	Related_drugs                 string
	Medical_condition_description string
	Rating                        float32
	No_of_reviews                 int
	Drug_link                     string
	Medical_condition_url         string
}

var Drugs []Drug

// Custom functions

func ReadDrugsFile() []Drug {
	jsonDrugFile, err := os.Open("drugs.json")

	if err != nil {
		fmt.Println(err)
	}

	var drugs []Drug
	byteValue, _ := ioutil.ReadAll(jsonDrugFile)
	json.Unmarshal(byteValue, &drugs)

	defer jsonDrugFile.Close()

	return drugs

}

func CreateNamesList(drugs []Drug) []string {

	drugNamesList := []string{}
	for i := 0; i < len(drugs); i++ {
		drugNamesList = append(drugNamesList, drugs[i].Drug_name)
	}

	return drugNamesList

}

func SearchByName(drugNames []string, searchTerm string) []string {
	searchResults := []string{}
	for i := 0; i < len(drugNames); i++ {
		if strings.Contains(drugNames[i], searchTerm) {
			searchResults = append(searchResults, drugNames[i])
		}
	}
	return searchResults
}

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

func SearchDrug(c echo.Context) error {

	searchTerm := c.Param("drugName")

	//var drugs []Drug = ReadDrugsFile()
	//var s = CreateNamesList(drugs)

	var s = CreateNamesList(Drugs)
	var res = SearchByName(s, searchTerm)

	jsonResponse := []byte{}
	jsonResponse, _ = json.Marshal(res)
	return c.JSONBlob(http.StatusOK, jsonResponse)

}
