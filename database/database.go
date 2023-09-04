package database

import (
	"context"

	"github.com/noodahl-org/libs/models"
	"gorm.io/gorm"
)

type Database interface {
	MigrateDomainModels(models ...*interface{}) error
	//shortened urls
	CreateSource(i *models.Source) (int64, error)
	GetOrCreateSource(i *models.Source) error
	FetchSource(q *models.Source) error
	//FetchSources(a *models.Source) ([]models.Source, error)
	DeleteSource(q *models.Source) error
	//rss
	CreateArticles(i ...models.Article) error
}

type PostgresDB struct {
	ctx    context.Context
	client *gorm.DB
}
