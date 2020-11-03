package repository

import "kumparan/domain/news/entity"

// Repository ...
type Repository interface {
	GetNews() ([]*entity.News, error)
	InsertNews(news *entity.News) error
}
