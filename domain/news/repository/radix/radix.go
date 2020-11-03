package radix

import (
	"encoding/json"
	"kumparan/domain/news/entity"
	"kumparan/domain/news/repository"
	"log"

	"github.com/mediocregopher/radix/v3"
)

const allNewsKey string = "news:all"

// Radix ...
type Radix struct {
	rdx       *radix.Pool
	mysqlRepo repository.Repository
}

// New ...
func New(rdx *radix.Pool, _mysqlRepo repository.Repository) *Radix {
	return &Radix{rdx: rdx, mysqlRepo: _mysqlRepo}
}

// GetNews ...
func (r *Radix) GetNews() ([]*entity.News, error) {
	var (
		b    []byte
		mn   = radix.MaybeNil{Rcv: &b}
		news []*entity.News
	)

	err := r.rdx.Do(radix.FlatCmd(&mn, "GET", allNewsKey))
	if err != nil {
		return nil, err
	} else if mn.Nil {
		news, err = r.mysqlRepo.GetNews()
		if err != nil {
			return nil, err
		} else if news == nil {
			return nil, nil
		}

		err = r.setNewsRadix(news)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(b, &news)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return news, nil
}

// InsertNews ...
func (r *Radix) InsertNews(news *entity.News) error {
	err := r.mysqlRepo.InsertNews(news)
	if err != nil {
		return err
	}

	err = r.rdx.Do(radix.FlatCmd(nil, "DEL", allNewsKey))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Radix) setNewsRadix(news []*entity.News) error {
	jNews, err := json.Marshal(news)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.rdx.Do(radix.FlatCmd(nil, "SET", allNewsKey, jNews))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
