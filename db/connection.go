package db

import (
	"database/sql",
	"fmt",
	"github.com/RobsonFeitosa/app-crud-go-lang/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		config.Host,
		config.Port,
		config.User,
		config.Pass,
		config.Database
	)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}