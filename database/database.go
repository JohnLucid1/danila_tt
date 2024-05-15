package database

import (
	"danila_first/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Environment struct {
	db_port int
	db_name string
	db_host string
	db_user string
	db_pass string
}

func NewEnv(db_port int, db_name, db_host, db_user, db_pass string) *Environment {
	p := Environment{
		db_port,
		db_name,
		db_host,
		db_user,
		db_pass,
	}
	return &p
}

func (env *Environment) Db_Init() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.db_host, env.db_port, env.db_user, env.db_pass, env.db_name) // просто создаем строку подключения
	// 					драйвер 
	db, err := sql.Open("postgres", connStr) // подлючаемся к бд
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAllCars(DB *sql.DB) ([]models.Car, error) {
	rows, err := DB.Query("select * from cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car

	for rows.Next() {
		var car models.Car

		if err := rows.Scan(
			&car.ID,
			&car.Name,
			&car.CarModelId,
		); err != nil {
			return nil, err
		}

		cars = append(cars, car)
	}

	return cars, nil
}

func GetCar(DB *sql.DB, id int) (models.Car, error) {
	var car models.Car
	err := DB.QueryRow("SELECT * FROM cars WHERE id= $1", id).Scan(
		&car.ID,
		&car.Name,
		&car.CarModelId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return car, err
		}
	}

	return car, nil
}

//
