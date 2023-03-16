package responses

import (
	"github.com/labstack/echo/v4"
)

// To change!
type UserResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    *echo.Map `json:"data"`
}
