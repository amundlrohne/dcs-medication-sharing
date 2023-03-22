package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/responses"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type GCPServiceAccountJSON struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func generateBearerToken() (string, error) {
	var s []byte = []byte(configs.EnvGCPServiceKey())
	credentials := GCPServiceAccountJSON{}
	json.Unmarshal(s, &credentials)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "medication-record-account@dcs-medication-sharing.iam.gserviceaccount.com",
		"sub": "medication-record-account@dcs-medication-sharing.iam.gserviceaccount.com",
		"aud": "https://healthcare.googleapis.com/",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 3600).Unix(),
	})

	token.Header["kid"] = credentials.PrivateKeyId

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(credentials.PrivateKey))
	if err != nil {
		fmt.Errorf("error parsing RSA private key: %v\n", err)
	}

	return token.SignedString(parsedKey)
}

func fhirGet(query string) (*http.Request, error) {
	prod := configs.EnvProduction()

	if prod == "true" {
		url := "https://healthcare.googleapis.com/v1/projects/dcs-medication-sharing/locations/europe-west4/datasets/dcs-medication-sharing-dataset/fhirStores/dcs-medication-sharing-fhir-store/fhir/" + query

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		token, err := generateBearerToken()
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+token)
		return req, nil

	} else {
		url := "http://" + configs.FHIR_URI() + ":8080/fhir/" + query

		return http.NewRequest("GET", url, nil)
	}

}

func fhirPost(query string, body []byte) (*http.Request, error) {
	prod := configs.EnvProduction()

	if prod == "true" {
		url := "https://healthcare.googleapis.com/v1/projects/dcs-medication-sharing/locations/europe-west4/datasets/dcs-medication-sharing-dataset/fhirStores/dcs-medication-sharing-fhir-store/fhir/" + query

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}

		token, err := generateBearerToken()
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
		req.Header.Set("Authorization", "Bearer "+token)

		return req, nil
	} else {
		url := "http://" + configs.FHIR_URI() + ":8080/fhir/" + query

		return http.NewRequest("POST", url, bytes.NewBuffer(body))
	}
}

func fhirDelete(query string) (*http.Request, error) {
	prod := configs.EnvProduction()

	if prod == "true" {
		url := "https://healthcare.googleapis.com/v1/projects/dcs-medication-sharing/locations/europe-west4/datasets/dcs-medication-sharing-dataset/fhirStores/dcs-medication-sharing-fhir-store/fhir/" + query

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return nil, err
		}

		token, err := generateBearerToken()
		if err != nil {
			return nil, err
		}

		req.Header.Set("Content-Type", "application/fhir+json; charset=UTF-8")
		req.Header.Set("Authorization", "Bearer "+token)

		return req, nil
	} else {
		url := "http://" + configs.FHIR_URI() + ":8080/fhir/" + query

		return http.NewRequest("DELETE", url, nil)
	}
}

func GetMedicationBundle(c echo.Context) error {
	var consent models.Consent
	// defer cancel()
	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	req, err := fhirGet("Bundle?identifier=" + consent.ConsentID)

	if err != nil {
		fmt.Errorf("error setting up medication-statement request: %v\n", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

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
	req, err := fhirGet("Bundle")

	if err != nil {
		fmt.Errorf("error setting up medication-statement request: %v\n", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//Convert the body to type string
	sb := string(body)

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": sb}})

}

func PostMedicationBundle(c echo.Context) error {
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
	req, err := fhirPost("Bundle", jsonValue)
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
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "Not able to post to FHIR", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	// Print the response status code and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", resp)

	return c.JSON(http.StatusCreated, responses.MedicationRecordResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": bundle}})
}

func DeleteMedicationBundle(c echo.Context) error {

	var consent models.Consent
	// defer cancel()
	if err := c.Bind(&consent); err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	req, err := fhirDelete("Bundle?identifier=" + consent.ConsentID)

	if err != nil {
		fmt.Errorf("error setting up medication-statement request: %v\n", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.MedicationRecordResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	return c.JSON(http.StatusOK, responses.MedicationRecordResponse{Status: http.StatusOK, Message: "Deleted successfully", Data: &echo.Map{"data": consent.ConsentID}})

}
