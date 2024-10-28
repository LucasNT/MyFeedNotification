package useCase

import (
	"context"
	"fmt"
	"net/url"
	"sort"

	"github.com/LucasNT/MyFeed/internal/feedParser"
)

func NewFeed(feedUrl *url.URL) *Feed {
	var f *Feed = new(Feed)
	f.FeedUrl = feedUrl
	return f
}

func (feed *Feed) GetFeed(ctx context.Context) error {
	parser := feedParser.NewParser()
	resp, err := feedParser.ParseFromURL(ctx, parser, feed.FeedUrl)
	if err != nil {
		return fmt.Errorf("Failed to get Feed from %s: %w", feed.FeedUrl.String(), err)
	}
	select {
	case <-ctx.Done():
		return fmt.Errorf("Context Canceled %s: %v", feed.FeedUrl.String(), ctx.Err())
	default:
	}
	sort.Sort(resp)
	select {
	case <-ctx.Done():
		return fmt.Errorf("Context Canceled %s: %v", feed.FeedUrl.String(), ctx.Err())
	default:
	}
	item := resp.Items[len(resp.Items)-1]
	feed.Link = item.Link
	feed.Message = item.Title
	feed.LastUpdate = item.PublishedParsed
	feed.Title = resp.Title
	return nil
}
