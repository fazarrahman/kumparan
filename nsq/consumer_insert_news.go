package nsq

import (
	"encoding/json"
	"kumparan/domain/news/entity"
	"kumparan/domain/news/repository"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

type NSQConsumer struct {
	newsRepo repository.Repository
}

func New(_repo repository.Repository) *NSQConsumer {
	return &NSQConsumer{newsRepo: _repo}
}

// News ...
type NewsInsertRequest struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

func (n *NSQConsumer) InitNSQ() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("news_insert", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var news NewsInsertRequest
		_ = json.Unmarshal(message.Body, &news)
		log.Printf("Message content : %v", news)

		log.Println("Inserting news data")
		err := n.newsRepo.InsertNews(&entity.News{
			Author: news.Author,
			Body:   news.Body,
		})

		if err != nil {
			log.Fatalln("Error when inserting news : ", err.Error())
		}

		log.Println("Data is successfully inserted")
		//wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect to nsq")
	}
	wg.Wait()
}

/*
func (n *NSQConsumer) ConsumerInsertNews(message *nsq.Message) error {
	{
		var news NewsInsertRequest
		_ = json.Unmarshal(message.Body, &news)
		log.Printf("Message content : %v", news)

		log.Println("Inserting news data")
		err := n.newsRepo.InsertNews(&entity.News{
			Author: news.Author,
			Body:   news.Body,
		})

		if err != nil {
			log.Fatalln("Error when inserting news : ", err.Error())
		}

		log.Println("Data is successfully inserted")
		//wg.Done()
		return nil
	}
}

*/
