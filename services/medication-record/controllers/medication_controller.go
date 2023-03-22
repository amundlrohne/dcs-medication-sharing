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

func GetMedicationRecord(c echo.Context) error {
	var consent models.Consent
	// defer cancel()

	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	url := "http://" + configs.FHIR_URI() + ":8080/fhir/MedicationStatement?identifier=" + consent.ConsentID

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

func GetAllMedicationRecords(c echo.Context) error {
	url := "http://" + configs.FHIR_URI() + ":8080/fhir/MedicationStatement?_pretty=true"

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
	// var medicationBundle models.MedicationBundle

	// if err := c.Bind(&medicationBundle); err != nil {
	// 	return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	// }

	// url := "http://" + configs.FHIR_URI() + ":8080/fhir/MedicationStatement?_format=json&_pretty=true"

	// bundle := &models.Bundle{
	// 	ID: medicationBundle.ConsentID,
	// 	Identifier: [1]models.Identifier{{Value: medicationBundle.ConsentID}},
	// 	Type: "history",
	// 	Entry: medicationBundle.MedicationRecords
	// }

	var medicationRecords models.MedicationRecords

	if err := c.Bind(&medicationRecords); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// var medicationStatements []models.MedicationStatement

	for i := 0; i < len(medicationRecords.Records); i++ {
		fmt.Println(medicationRecords.Records[i])
		// ms := &models.MedicationStatement{
		// 	ResourceType:              "MedicationStatement",
		// 	ID:                        medicationRecord.ConsentID,
		// 	Subject:                   models.Subject{Display: medicationRecord.Name},
		// 	Status:                    medicationRecord.Status,
		// 	MedicationCodeableConcept: models.MedicationCodeableConcept{Text: medicationRecord.Medication},
		// 	Note:                      [1]models.Note{{Text: medicationRecord.Note}},
		// 	EffectiveDateTime:         medicationRecord.EffectiveDateTime,
		// 	Dosage:                    [1]models.Dosage{{Sequence: medicationRecord.DosageSequence, Text: medicationRecord.DosageNote}},
		// 	Identifier:                [1]models.Identifier{{Value: medicationRecord.ConsentID}},
		// }

		// medicationStatements = append(medicationStatements, ms)
	}

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": "hey"}})

}

func PostMedicationRecord(c echo.Context) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var medicationRecord models.MedicationRecord
	// defer cancel()

	if err := c.Bind(&medicationRecord); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	// Define the FHIR server endpoint
	url := "http://" + configs.FHIR_URI() + ":8080/fhir/MedicationStatement?_format=json&_pretty=true"

	fmt.Println("URL :: " + url)

	mr := &models.MedicationStatement{
		ResourceType:              "MedicationStatement",
		ID:                        medicationRecord.ConsentID,
		Subject:                   models.Subject{Display: medicationRecord.Name},
		Status:                    medicationRecord.Status,
		MedicationCodeableConcept: models.MedicationCodeableConcept{Text: medicationRecord.Medication},
		Note:                      [1]models.Note{{Text: medicationRecord.Note}},
		EffectiveDateTime:         medicationRecord.EffectiveDateTime,
		Dosage:                    [1]models.Dosage{{Sequence: medicationRecord.DosageSequence, Text: medicationRecord.DosageNote}},
		Identifier:                [1]models.Identifier{{Value: medicationRecord.ConsentID}},
	}

	// Encode the JSON object as a byte slice
	jsonValue, _ := json.Marshal(mr)

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

	fmt.Println("HEREEEE...")
	return c.JSON(http.StatusCreated, responses.MedicationRecordResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": medicationRecord.ConsentID}})

}

func DeleteMedicationRecord(c echo.Context) error {
	var consent models.Consent
	// defer cancel()

	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	url := "http://" + configs.FHIR_URI() + ":8080/fhir/MedicationStatement?identifier=" + consent.ConsentID

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
