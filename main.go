package main

import (

	"net/http"

	"github.com/labstack/echo/v4"
)



func main() {
	app := echo.New()	

	app.GET("/", func (c echo.Context)  error{
		return c.String(http.StatusOK, "hello world")
	})

	app.Logger.Fatal(app.Start(":8080"))
}

// https://github.com/gofiber/fiber
// go get github.com/lib/pq => Работ