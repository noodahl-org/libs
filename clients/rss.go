package clients

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/noodahl-org/libs/config"
	"github.com/noodahl-org/libs/database"
	"github.com/noodahl-org/libs/models"
)

type Feed interface {
	FetchURL(ctx context.Context, c chan []models.Article, url string, sourceID uuid.UUID)
	FetchFeedList(ctx context.Context, feedSources []models.FeedSources)
}

type RSSClient struct {
	db     database.Database
	conf   *config.Conf
	parser *gofeed.Parser
}

func NewRSSClient(conf *config.Conf, db database.Database) Feed {
	return RSSClient{
		conf:   conf,
		db:     db,
		parser: gofeed.NewParser(),
	}
}

func (r RSSClient) FetchFeedList(ctx context.Context, feedSources []models.FeedSources) {
	var ch chan []models.Article
	var totalEPs int
	defer close(ch)

	for _, fs := range feedSources {
		for _, endpoint := range fs.Endpoints {
			totalEPs += len(fs.Endpoints)
			source := &models.Source{
				URL:  endpoint,
				Tags: fs.Tags,
			}
			err := r.db.GetOrCreateSource(source)
			if err != nil {
				log.Panic(err)
			}
			go r.FetchURL(ctx, ch, endpoint, source.StorageBase.ID)
		}
	}

	for i := 0; i < totalEPs; i++ {
		feed := <-ch
		if len(feed) > 0 {
			num, err := r.db.CreateArticles(feed)
			log.Printf("%v articles created", num)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

func (r RSSClient) FetchURL(ctx context.Context, c chan []models.Article, url string, sourceID uuid.UUID) {
	out := []models.Article{}
	feed, err := r.parser.ParseURLWithContext(url, ctx)
	if err != nil {
		log.Printf("error: %s", err)
		c <- []models.Article{}
	}

	for _, item := range feed.Items {
		article, err := models.ArticleFromFeedItem(item)
		article.SourceID = sourceID
		if err != nil {
			log.Printf("error: %s", err)
			c <- []models.Article{}
		}
		out = append(out, *article)
	}
	c <- out
}
