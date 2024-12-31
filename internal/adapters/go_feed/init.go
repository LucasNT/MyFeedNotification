package gofeed

import (
	"context"
	"fmt"
	"net/url"
	"sort"

	"github.com/LucasNT/MyFeed/internal/entities"
	"github.com/mmcdole/gofeed"
)

type GoFeed struct {
	feedParser *gofeed.Parser
}

func NewGoFeed() GoFeed {
	ret := GoFeed{
		feedParser: gofeed.NewParser(),
	}
	return ret
}

func (g GoFeed) GetFeed(ctx context.Context, link *url.URL) (entities.Feed, error) {
	var ret entities.Feed
	resp, err := g.feedParser.ParseURLWithContext(link.String(), ctx)
	if err != nil {
		return entities.Feed{}, fmt.Errorf("Failed to parse rss feed: %w", err)
	}
	sort.Sort(resp)
	select {
	case <-ctx.Done():
		return entities.Feed{}, fmt.Errorf("Context Canceled %s: %w", link.String(), ctx.Err())
	default:
	}
	item := resp.Items[len(resp.Items)-1]
	auxLink, err := url.Parse(item.Link)
	if err != nil {
		return entities.Feed{}, fmt.Errorf("Failed to parse url %s, %w", item.Link, err)
	}
	ret = entities.NewFeed(resp.Title, item.Title, *item.PublishedParsed, auxLink, entities.NORMAL_LEVEL)
	return ret, nil
}
