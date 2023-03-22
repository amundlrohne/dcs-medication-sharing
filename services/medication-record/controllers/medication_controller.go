package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/responses"
	"github.com/labstack/echo/v4"
)

func GetMedicationBundle(c echo.Context) error {
	var consent models.Consent
	// defer cancel()
	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	url := "http://" + configs.FHIR_URI() + ":8080/fhir/Bundle?identifier=" + consent.ConsentID

	resp, err := http.Get(url)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//Convert the body to type string
	sb := string(body)

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": sb}})
}

func GetAllMedicationBundles(c echo.Context) error {
	url := "http://" + configs.FHIR_URI() + ":8080/fhir/Bundle?_pretty=true"

	resp, err := http.Get(url)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//Convert the body to type string
	sb := string(body)

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": sb}})

}

func PostMedicaitonBundle(c echo.Context) error {

	url := "http://" + configs.FHIR_URI() + ":8080/fhir/Bundle?_format=json&_pretty=true"

	var medicationRecords models.MedicationRecords

	if err := c.Bind(&medicationRecords); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	var medicationStatements []models.MedicationStatement
	var resources []models.Resource

	for i := 0; i < len(medicationRecords.Records); i++ {
		fmt.Println(medicationRecords.Records[i])
		ms := models.MedicationStatement{
			ResourceType:              "MedicationStatement",
			ID:                        medicationRecords.Records[i].ConsentID,
			Subject:                   models.Subject{Display: medicationRecords.Records[i].Name},
			Status:                    medicationRecords.Records[i].Status,
			MedicationCodeableConcept: models.MedicationCodeableConcept{Text: medicationRecords.Records[i].Medication},
			Note:                      [1]models.Note{{Text: medicationRecords.Records[i].Note}},
			EffectiveDateTime:         medicationRecords.Records[i].EffectiveDateTime,
			Dosage:                    [1]models.Dosage{{Sequence: medicationRecords.Records[i].DosageSequence, Text: medicationRecords.Records[i].DosageNote}},
			Identifier:                [1]models.Identifier{{Value: medicationRecords.Records[i].ConsentID}},
		}

		medicationStatements = append(medicationStatements, ms)
	}

	for i := 0; i < len(medicationStatements); i++ {
		rs := models.Resource{
			MedicationStatement: medicationStatements[i],
		}

		resources = append(resources, rs)
	}

	bundle := &models.Bundle{
		ResourceType: "Bundle",
		ID:           medicationRecords.Records[0].ConsentID,
		Identifier:   [1]models.Identifier{{Value: medicationRecords.Records[0].ConsentID}},
		Type:         "collection",
		Entry:        resources,
	}

	jsonValue, _ := json.Marshal(bundle)

	// Create a new HTTP request with the POST method
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// Set the appropiate request headers
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
	req.Header.Set("Accept", "application/fhir+json;q=1.0, application/json+fhir;q=0.9")
	req.Header.Set("User-Agent", "HAPI-FHIR/6.4.0 (FHIR Client; FHIR 4.0.1/R4; apache)")

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	// Print the response status code and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", resp)

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": bundle}})

}

func DeleteMedicationBundle(c echo.Context) error {
	var consent models.Consent
	// defer cancel()

	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	url := "http://" + configs.FHIR_URI() + ":8080/fhir/Bundle?identifier=" + consent.ConsentID

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "Delted successfully", Data: &echo.Map{"data": ""}})
}
