package database

import (
	"github.com/noodahl-org/libs/models"
)

type Database interface {
	MigrateDomainModels(models ...*interface{}) error
	CreateSource(i *models.Source) (int64, error)
	GetOrCreateSource(i *models.Source) error
	UpdateSource(q *models.Source) (int64, error)
	FetchSource(q *models.Source) error
	DeleteSource(q *models.Source) error
	CreateArticles(i []models.Article) (int64, error)
	FetchArticles(q *models.Article) ([]models.Article, error)
	FetchArticle(q *models.Article) error
	FetchArticlesSummary(q *models.Article) ([]models.Article, error)
	FetchArticlesSummaryLimit(a *models.Article, limit int) ([]models.Article, error)
	CreateWebAPIRequestStats(i *models.APIRequestStats) error
	FetchDistinctSources() ([]string, error)
	CreateScrape(i *models.Scrape) (int64, error)
	CreateScrapes(i []models.Scrape) (int64, error)
	FetchScrape(q *models.Scrape) error
	FetchScrapes(q *models.Scrape, limit int) ([]models.Scrape, error)
	UpdateScrape(q *models.Scrape) (int64, error)
	LogError(caller string, err error) error
}
