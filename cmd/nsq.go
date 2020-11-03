package main

import (
	db "kumparan/config/mysql"
	rdx "kumparan/config/radix"
	news_mysql "kumparan/domain/news/repository/mysql"
	rdx_repo "kumparan/domain/news/repository/radix"
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

	radixInit, err := rdx.New()
	if err != nil {
		log.Println(err)
	}
	log.Println("Redis has been successfully initialized")

	newsMysqlRepo := news_mysql.New(dbClient)
	rdxRepository := rdx_repo.New(radixInit, newsMysqlRepo)
	log.Println("Repositories are successfully initialized")

	kumparanNsq.New(rdxRepository).InitNSQ()
}
