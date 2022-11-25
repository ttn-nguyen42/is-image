package repositories

import "is-image/core/models/entities"

/*
 * Interfaces (ports)
 */
type IResultCacheRepository interface {
	Find(id string) (*entities.ResultCache, error)
	Add(data *entities.ResultCache) error
}

type ResultCacheRepository struct {
}

func NewResultCacheRepository() *ResultCacheRepository {
	return &ResultCacheRepository{}
}

func (p *ResultCacheRepository) Find(id string) (*entities.ResultCache, error) {
	return &entities.ResultCache{}, nil
}

func (p *ResultCacheRepository) Add(data *entities.ResultCache) error {
	return nil
}
