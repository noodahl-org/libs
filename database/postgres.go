package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/noodahl-org/libs/config"
	"github.com/noodahl-org/libs/models"
	"github.com/noodahl-org/libs/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PgErrorNotFound error
type PgErrorConflict error

type PostgresDB struct {
	ctx    context.Context
	client *gorm.DB
}

func (d PostgresDB) MigrateDomainModels(models ...*interface{}) error {
	for _, model := range models {
		if err := d.client.WithContext(d.ctx).AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}

func NewPostgresDB(conf *config.Conf) (Database, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%v database=%s",
		conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBPort, conf.DBName)

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return PostgresDB{
		client: client,
	}, err

}

func (d PostgresDB) CreateArticles(articles []models.Article) (int64, error) {
	result := d.client.Clauses(clause.OnConflict{DoNothing: true}).Create(articles)
	return result.RowsAffected, result.Error
}

func (d PostgresDB) CreateArticle(i *models.Article) (int64, error) {
	result := d.client.Clauses(clause.OnConflict{DoNothing: true}).Create(i)
	return result.RowsAffected, result.Error
}

func (d PostgresDB) CreateSource(i *models.Source) (int64, error) {
	result := d.client.Clauses(clause.OnConflict{DoNothing: true}).Create(i)
	return result.RowsAffected, result.Error
}

func (d PostgresDB) GetOrCreateSource(i *models.Source) error {
	var err error
	url := i.URL
	err = d.FetchSource(i)
	if err != nil {
		switch {
		case err == gorm.ErrRecordNotFound:
			i.URL = url
			i.Domain = utils.Domain(url)
			i.StorageBase = models.StorageBase{
				ID: uuid.New(),
			}
			err = nil
			_, err = d.CreateSource(i)
		}
	}
	return err
}

func (d PostgresDB) FetchSource(q *models.Source) error {
	result := d.client.First(q, q)
	return result.Error
}

func (d PostgresDB) FetchArticlesSummary(q *models.Article) ([]models.Article, error) {
	out := []models.Article{}
	result := d.client.Omit("raw, content").Order("published DESC").Find(&out, q)
	return out, result.Error
}

func (d PostgresDB) FetchArticlesSummaryLimit(q *models.Article, limit int) ([]models.Article, error) {
	out := []models.Article{}
	result := d.client.Omit("raw, content").Limit(limit).Find(&out, q)
	return out, result.Error
}

func (d PostgresDB) FetchArticles(q *models.Article) ([]models.Article, error) {
	out := []models.Article{}
	result := d.client.Order("published DESC").Find(&out, q)
	return out, result.Error
}

func (d PostgresDB) DeleteSource(q *models.Source) error {
	result := d.client.WithContext(d.ctx).Where("id = ?", q.StorageBase.ID).Delete(&models.Source{})
	return result.Error
}

func (d PostgresDB) CreateWebAPIRequestStats(i *models.APIRequestStats) error {
	result := d.client.Clauses(clause.OnConflict{DoNothing: true}).Create(i)
	return result.Error
}
