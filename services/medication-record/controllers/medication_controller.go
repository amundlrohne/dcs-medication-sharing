package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//var medicationRecordCollection *mongo.Collection = configs.GetCollection(configs.DB, "medicationRecord")

var dbDummy = [10]string{"Hello", "world", "How", "are", "you"}

func GetMedication(c echo.Context) error {
	idStr := c.Param("id")

	//Convert id to int
	// id, err := strconv.Atoi(idStr)

	// if err != nil {
	// 	fmt.Println("Error during conversion")

	// }

	return c.String(http.StatusOK, "GET medication-record/"+idStr)
	//return c.String(http.StatusOK, "GET medication-record/"+idStr+" DUMMY DB DATA: "+dbDummy[id])
}

func PostMedication(c echo.Context) error {
	return c.String(http.StatusOK, "POST medication-record/ testing...")

	//Post to db testing
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// newMR := models.MedicationRecord{
	// 	Id:   primitive.NewObjectID(),
	// 	Name: "MedicationNameTest",
	// }
	// result, err := medicationRecordCollection.InsertOne(ctx, newMR)

	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	// }

	// return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})

}

func PutMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "PUT medication-record/"+id)
}

func DeleteMedication(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "DELETE medication-record/"+id)
}
