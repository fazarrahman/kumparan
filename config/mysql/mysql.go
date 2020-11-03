package mysql

import (
	"kumparan/helper"
	"log"

	"github.com/jmoiron/sqlx"
)

func New() (*sqlx.DB, error) {
	var password string = helper.GetEnv("DB_PASSWORD", "")
	var server string = helper.GetEnv("DB_SERVER", "")
	var dbName string = helper.GetEnv("DB_DATABASE_NAME", "")
	var username string = helper.GetEnv("DB_USERNAME", "")

	db, err := sqlx.Connect("mysql", username+":"+password+"@("+server+")/"+dbName)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return db, nil
}
