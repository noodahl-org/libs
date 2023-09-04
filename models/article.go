package models

import (
	"fmt"
	"net/url"

	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"jaytaylor.com/html2text"

	"github.com/noodahl-org/libs/utils"
)

type Links []string
type Tags []string

type Article struct {
	StorageBase
	SourceID  uuid.UUID  `json:"source_id"`
	Domain    string     `json:"domain"`
	Hash      string     `json:"hash" gorm:"index:idx_article_hash,unique"`
	Title     string     `json:"title"`
	Raw       string     `json:"raw"`
	Content   string     `json:"content"`
	Published *time.Time `json:"published"`
	Link      string     `json:"link"`
	Links     Links      `json:"links" gorm:"type:text"`
	Tags      Tags       `json:"tags" gorm:"type:text"`
}

func ArticleFromFeedItem(item *gofeed.Item) (*Article, error) {
	var err error
	u, err := url.Parse(item.Link)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	if item.PublishedParsed == nil {
		item.PublishedParsed = &now
	}
	article := Article{
		Title:     item.Title,
		Raw:       item.Description,
		Link:      item.Link,
		Published: item.PublishedParsed,
		Domain:    u.Host,
		StorageBase: StorageBase{
			ID:        uuid.New(),
			CreatedAt: time.Now().Unix(),
		},
		Hash:  utils.Hash(fmt.Sprintf("%s%s%s", u.Host, item.Title, item.Link)),
		Links: []string{},
		Tags:  []string{},
	}
	article.Content, err = html2text.FromString(item.Description, html2text.Options{
		PrettyTables: false,
		TextOnly:     true,
		OmitLinks:    true,
	})
	if err != nil {
		return nil, err
	}

	return &article, nil
}
