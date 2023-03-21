package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

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
var DrugNames []string

// Parse json drug library into a js drug array
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

// Return all drugs as JSON
func AllDrugNames(c echo.Context) error {

	jsonResponse := []byte{}
	jsonResponse, _ = json.Marshal(DrugNames)
	return c.JSONBlob(http.StatusOK, jsonResponse)
}

// Create a js string array containing all drug names
func CreateNamesList(drugs []Drug) []string {

	drugNamesList := []string{}
	for i := 0; i < len(drugs); i++ {
		drugNamesList = append(drugNamesList, drugs[i].Drug_name)
	}

	return drugNamesList

}

// Search based on the seatchTerm through all the drug names string array
// and return a string array as a result
func SearchByName(drugNames []string, searchTerm string) []string {
	var searchTerm2 = strings.Title(searchTerm)
	searchResults := []string{}
	for i := 0; i < len(drugNames); i++ {
		if strings.Contains(drugNames[i], searchTerm) || (strings.Contains(drugNames[i], searchTerm2)) {
			searchResults = append(searchResults, drugNames[i])
		}
	}
	return searchResults
}

// Return a json based on the SearchByName function
func SearchDrug(c echo.Context) error {

	searchTerm := c.Param("drugName")

	var res = SearchByName(DrugNames, searchTerm)

	jsonResponse := []byte{}
	jsonResponse, _ = json.Marshal(res)
	return c.JSONBlob(http.StatusOK, jsonResponse)

}