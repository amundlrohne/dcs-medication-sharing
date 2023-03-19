package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/responses"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var providerCollection *mongo.Collection = configs.GetCollection(configs.DB, "providers")
var validate = validator.New()

func CreateProvider(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var provider models.Provider
	defer cancel()

	//validate the request body
	if err := c.Bind(&provider); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProviderResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&provider); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.ProviderResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newProvider := models.Provider{
		Name:     provider.Name,
		Country:  provider.Country,
		Password: HashPassword(provider.Password),
		Username: provider.Username,
	}

	result, err := providerCollection.InsertOne(ctx, newProvider)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.ProviderResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetProvider(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	providerID := c.Param("provider_id")
	var provider models.Provider
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(providerID)

	err := providerCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&provider)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": provider}})
}

func GetAllProviders(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var providers []models.Provider
	defer cancel()

	results, err := providerCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleProvider models.Provider
		if err = results.Decode(&singleProvider); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		providers = append(providers, singleProvider)
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": providers}})
}

func VerifyUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	var provider models.Provider
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProviderResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	err := providerCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&provider)

	check, msg := ComparePasswords(user.Password, provider.Password)

	if !check {
		return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: msg, Data: &echo.Map{"data": nil}})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	writeAuthCookie(c, user)
	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": provider}})
}

func writeAuthCookie(c echo.Context, u models.User) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = u.Username
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func ReadAuthCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": cookie.Value}})
}

func DeleteAuthCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(0)
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "Cookie Deleted...", Data: &echo.Map{"data": ""}})
}

func ComparePasswords(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "login or passowrd is incorrect"
		check = false
	}

	return check, msg
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "Error"
	}

	return string(bytes)
}
