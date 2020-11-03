package main

import (
	db "kumparan/config/mysql"
	news_mysql "kumparan/domain/news/repository/mysql"
	kumparanNsq "kumparan/nsq"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func nsqCmd() {
	dbClient, err := db.New()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database is successfully initialized")

	newsMysqlRepo := news_mysql.New(dbClient)
	log.Println("News repository is successfully initialized")

	kumparanNsq.New(newsMysqlRepo).InitNSQ()
}
