package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kumparan/domain/news/repository"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/nsqio/go-nsq"
)

// Rest ...
type Rest struct {
	newsRepo repository.Repository
	producer *nsq.Producer
}

// New ...
func New(_repo repository.Repository, _producer *nsq.Producer) *Rest {
	return &Rest{newsRepo: _repo, producer: _producer}
}

// News ...
type News struct {
	Author  string    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// PostNewsRequest ...
type PostNewsRequest struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

// Register ...
func (r *Rest) Register(router *mux.Router) {
	router.HandleFunc("/news", r.GetNews).Methods("GET")
	router.HandleFunc("/news", r.PostNews).Methods("POST")

}

// GetNews ...
func (rest *Rest) GetNews(w http.ResponseWriter, r *http.Request) {
	var news []*News

	newsEnt, err := rest.newsRepo.GetNews()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if newsEnt == nil {
		w.Write([]byte("Data not found"))
		return
	}

	for _, n := range newsEnt {
		news = append(news, &News{
			Author:  n.Author,
			Body:    n.Body,
			Created: n.Created,
		})
	}

	b, err := json.Marshal(news)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

// PostNews ...
func (rest *Rest) PostNews(w http.ResponseWriter, r *http.Request) {
	var news PostNewsRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &news)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshal")
		return
	}

	if strings.Trim(news.Author, " ") == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Authors cannot be blank")
		return
	} else if strings.Trim(news.Body, " ") == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Body cannot be blank")
		return
	}

	log.Println("Publishing news : ")
	log.Println(news)
	err = rest.producer.Publish("news_insert", reqBody)
	if err != nil {
		log.Panic("Could not connect to nsq")
	}

	log.Println("Successfully publish the news data")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(news)
}
