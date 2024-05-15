package main

import (
	"danila_first/database"
	"fmt"
	"log"
	"os"
	"strconv"

	"danila_first/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load() // Загружаем env файл
	if err != nil {
		log.Fatalln(err)
	}

	app := echo.New() // Создаем инстанцию приложения (от которого идет почти все)
	PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	env := database.NewEnv(
		PORT,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
	)

	db, err := env.Db_Init()
	fmt.Println(db.Ping())
	if err != nil {
		log.Fatalln(err)
	}

	app.GET("/cars", handlers.GetAllCars(db)) // обычный роут с коллбеком
	app.GET("/cars/:id", handlers.GetCar(db)) // динамический роут с коллбеком
	app.Start(":8080")
}

// go get github.com/labstack/echo/v4 => сервер
// go get github.com/lib/pq => драйвер датабазы чтобы писать чистые запросы на sql
