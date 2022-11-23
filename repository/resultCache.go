package repository

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
 * MongoDB repository of a cache database
 */
type ResultCacheRepository struct {
	ctx    context.Context
	client *mongo.Client
}

var lock = &sync.Mutex{}
var instance *ResultCacheRepository

/*
 * Get singleton
 */
func GetResultCacheRepositoryInstance() *ResultCacheRepository {
	return instance
}

func NewResultCacheRepository(opts *options.ClientOptions, ctx context.Context) (*ResultCacheRepository, error) {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			client, err := mongo.Connect(ctx, opts)
			defer func() {
				err = client.Disconnect(ctx)
				if err != nil {
					log.Printf("Failed to disconnect MongoDB client")
				}
			}()
			if err != nil {
				return nil, err
			}
			instance = &ResultCacheRepository{
				client: client,
				ctx:    ctx,
			}
		}
	}
	return instance, nil
}

/*
 * Return the MongoDB client
 */
func (p *ResultCacheRepository) Client() (*mongo.Client, error) {
	if p.client == nil {
		return nil, fmt.Errorf("MongoDB client is nil")
	}
	return p.client, nil
}

func (p *ResultCacheRepository) Context() (context.Context, error) {
	if p.ctx == nil {
		return nil, fmt.Errorf("MongoDB context is nil")
	}
	return p.ctx, nil
}
