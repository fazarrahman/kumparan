package main

import (
	"fmt"
	db "kumparan/config/mysql"
	news_mysql "kumparan/domain/news/repository/mysql"
	"kumparan/helper"

	"kumparan/rest/external"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nsqio/go-nsq"
)

func app() {
	dbClient, err := db.New()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database is successfully initialized")

	newsMysqlRepo := news_mysql.New(dbClient)
	log.Println("News repository is successfully initialized")

	// producer
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatalln("Failed to initialize nsq producer")
		return
	}

	log.Println("NSQ Producer is successfully initialized")

	router := mux.NewRouter()
	external.New(newsMysqlRepo, producer).Register(router)

	http.Handle("/", router)
	fmt.Println("Connected to port " + helper.GetEnv("APP_PORT", ""))
	log.Fatal(http.ListenAndServe(":"+helper.GetEnv("APP_PORT", ""), router))
}
