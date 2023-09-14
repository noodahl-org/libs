package clients

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/noodahl-org/libs/config"
	"github.com/noodahl-org/libs/database"
	"github.com/noodahl-org/libs/models"
)

type Feed interface {
	FetchURL(ctx context.Context, c chan []models.Article, url string, sourceID uuid.UUID)
	FetchFeedList(ctx context.Context, feedSources []string)
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

func (r RSSClient) FetchFeedList(ctx context.Context, feedSources []string) {
	ch := make(chan []models.Article, len(feedSources))
	defer close(ch)

	for _, url := range feedSources {
		source := &models.Source{
			URL: url,
		}
		err := r.db.GetOrCreateSource(source)
		if err != nil {
			r.db.LogError(fmt.Sprintf("RSSClient.FetchFeedList: %v", url), err)
			continue
		}
		go r.FetchURL(ctx, ch, url, source.StorageBase.ID)

	}

	for i := 0; i < len(feedSources); i++ {
		feed := <-ch
		if len(feed) > 0 {
			num, err := r.db.CreateArticles(feed)
			if err != nil {
				log.Panic(err)
			}
			log.Printf("%v articles created", num)
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
