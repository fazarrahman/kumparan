package repository

import "kumparan/domain/news/entity"

type Repository interface {
	GetNews() ([]*entity.News, error)
	InsertNews(news *entity.News) error
}
