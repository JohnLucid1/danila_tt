package handlers

import (
	"danila_first/database"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func GetAllCars(DB sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: write the request
	}
}