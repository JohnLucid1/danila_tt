package database

import (
	"database/sql"
	"fmt"
)

type Environment struct {
	db_port int
	db_name string
	db_host string
	db_user string
	db_pass string
}

func (env *Environment) Db_Init() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.db_host, env.db_port, env.db_user, env.db_pass, env.db_name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//
