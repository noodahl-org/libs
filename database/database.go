package database

import (
	"github.com/noodahl-org/libs/models"
)

type Database interface {
	MigrateDomainModels(models ...*interface{}) error
	CreateSource(i *models.Source) (int64, error)
	GetOrCreateSource(i *models.Source) error
	FetchSource(q *models.Source) error
	DeleteSource(q *models.Source) error
	CreateArticles(i ...models.Article) error
	CreateWebAPIRequestStats(i *models.APIRequestStats) error
}
