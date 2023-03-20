package main

import (
	"fmt"

	"github.com/amundlrohne/dcs-medication-sharing/services/medication-record/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	fmt.Println(Add(9, 1))

	routes.MedicationRoute(e)

	e.Logger.Fatal(e.Start(":8080"))

}

// Testing dummy
func Add(x, y int) (res int) {
	return x + y
}
