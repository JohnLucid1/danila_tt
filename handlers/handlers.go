package handlers

import (
	"danila_first/database"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCars(DB *sql.DB) echo.HandlerFunc { // создаем wrapper функцию чтобы она принимала указатель к датабазе
	return func(c echo.Context) error { // возвращаем хэндлер функию
		cars, err := database.GetAllCars(DB) // берем данные из бд
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, cars) // возвращаем их в json формате с http статусом 200
	}
}

func GetCar(DB *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		car, err := database.GetCar(DB, id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, car)
	}
}
