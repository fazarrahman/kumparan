package mysql

import (
	"database/sql"
	"kumparan/domain/news/entity"
	"log"

	"github.com/jmoiron/sqlx"
)

// Mysqldb ...
type Mysqldb struct {
	db *sqlx.DB
}

// New ...
func New(_db *sqlx.DB) *Mysqldb {
	return &Mysqldb{db: _db}
}

// GetNews ...
func (m *Mysqldb) GetNews() ([]*entity.News, error) {
	var news []*entity.News
	err := m.db.Select(&news,
		`select id, author, body, created
		from news
		order by created desc`)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return news, nil
}

// InsertNews ...
func (m *Mysqldb) InsertNews(news *entity.News) error {
	_, err := m.db.Exec("insert into news (author, body) values(?,?)", news.Author, news.Body)
	if err != nil {
		log.Println(err)
		return nil
	}

	return err
}
