package repositories

import (
	"is-image/core/models/entities"
	"is-image/db"
)

/*
 * Interfaces (ports)
 */
type IResultCacheRepository interface {
	Find(id string) (*entities.ResultCache, error)
	Add(data *entities.ResultCache) error
}

type ResultCacheRepository struct {
	db *db.MongoClient
}

func NewResultCacheRepository(db *db.MongoClient) *ResultCacheRepository {
	return &ResultCacheRepository{
		db: db,
	}
}

func (p *ResultCacheRepository) Find(id string) (*entities.ResultCache, error) {
	return &entities.ResultCache{}, nil
}

func (p *ResultCacheRepository) Add(data *entities.ResultCache) error {
	return nil
}
