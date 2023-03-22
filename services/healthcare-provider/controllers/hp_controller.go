package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/configs"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/models"
	"github.com/amundlrohne/dcs-medication-sharing/services/healthcare-provider/responses"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var providerCollection *mongo.Collection = configs.GetCollection(configs.DB, "providers")
var jwtCollection *mongo.Collection = configs.GetCollection(configs.DB, "jwts")
var validate = validator.New()

var jwtKey = []byte(configs.JWTSecretKey())

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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
		ID:       primitive.NewObjectID(),
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
	providerID := c.Param("id")
	var provider models.Provider
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(providerID)

	err := providerCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&provider)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": provider}})
}

// Search provider by name and return username
func GetProviderByName(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	searchName := c.Param("name")
	var provider models.Provider
	defer cancel()

	err := providerCollection.FindOne(ctx, bson.M{"name": searchName}).Decode(&provider)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": "Error finding HP based on name"}})
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": provider.ID}})
}

func DeleteProvider(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	resp, err := providerCollection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": "Error finding HP based on name"}})
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": resp}})
}

func GetProviderFromToken(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	token := c.Param("token")
	var provider models.Provider
	var jwtModel models.JWT
	defer cancel()

	mongo_err := jwtCollection.FindOne(ctx, bson.M{"token": token}).Decode(&jwtModel)
	if mongo_err != nil {
		return c.JSON(http.StatusUnauthorized, responses.ProviderResponse{Status: http.StatusUnauthorized, Message: "unauthorized", Data: &echo.Map{"data": mongo_err.Error()}})
	}

	err := providerCollection.FindOne(ctx, bson.M{"username": jwtModel.Username}).Decode(&provider)

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

	// Set expiration time to 30 days
	expirationTime := time.Now().Add(720 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	newJWT := models.JWT{
		Username: provider.Username,
		Token:    tokenString,
	}

	result, err := jwtCollection.InsertOne(ctx, newJWT)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	writeAuthCookie(c, user, tokenString, expirationTime)
	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": provider, "token_data": result}})
}

func writeAuthCookie(c echo.Context, u models.User, tokenString string, exp time.Time) error {

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = exp
	c.SetCookie(cookie)

	fmt.Println(c.Path())
	return nil
}

func ReadAuthCookie(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var jwtModel models.JWT
	defer cancel()

	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "no-cookie", Data: &echo.Map{"data": err.Error()}})
	}

	tknStr := cookie.Value

	mongo_err := jwtCollection.FindOne(ctx, bson.M{"token": tknStr}).Decode(&jwtModel)
	if mongo_err != nil {
		return c.JSON(http.StatusUnauthorized, responses.ProviderResponse{Status: http.StatusUnauthorized, Message: "unauthorized", Data: &echo.Map{"data": mongo_err.Error()}})
	}

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtModel.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.JSON(http.StatusUnauthorized, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "unauthorized", Data: &echo.Map{"data": err.Error()}})
		}
		return c.JSON(http.StatusBadRequest, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	if !tkn.Valid {
		return c.JSON(http.StatusUnauthorized, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "unauthorized", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": cookie.Value}})
}

func DeleteAuthCookie(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.ProviderResponse{Status: http.StatusInternalServerError, Message: "no-cookie", Data: &echo.Map{"data": err.Error()}})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := jwtCollection.DeleteOne(ctx, bson.M{"token": cookie.Value})
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.ProviderResponse{Status: http.StatusBadRequest, Message: "Not deleted ...", Data: &echo.Map{"data": err.Error()}})
	}

	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name:    "token",
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour), // Set expires for older versions of IE
		Path:    "/health-provider/verify",
	})

	return c.JSON(http.StatusOK, responses.ProviderResponse{Status: http.StatusOK, Message: "Cookie Deleted...", Data: &echo.Map{"data": result}})
}

func ComparePasswords(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "invalid"
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
