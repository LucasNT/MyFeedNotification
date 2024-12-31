package mocks

import (
	"context"
	"fmt"
	"net/url"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
	"github.com/LucasNT/MyFeed/internal/entities"
)

type GetFeed struct {
	feeds map[string]entities.Feed
}

func NewFeedGetter(feed entities.Feed, link string) *GetFeed {
	ret := new(GetFeed)
	ret.feeds = make(map[string]entities.Feed)
	ret.AppendFeed(feed, link)
	return ret
}

func (g *GetFeed) AppendFeed(feed entities.Feed, link string) error {
	urlParsed, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("Failed to add feed to getFeed Mock %w", err)
	}
	g.feeds[urlParsed.String()] = feed
	return nil
}

func (g *GetFeed) GetFeed(ctx context.Context, link *url.URL) (entities.Feed, error) {
	urlString := link.String()
	ret, ok := g.feeds[urlString]
	if !ok {
		err := fmt.Errorf("%w, not found in map", interfaces.ErrNofeedToReturn)
		return entities.Feed{}, err
	}
	return ret, nil
}
