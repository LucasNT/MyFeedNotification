package feedParser

import (
	"context"
	"fmt"
	"net/url"
	"sort"

	"github.com/mmcdole/gofeed"
)

func ParseFromURL(ctx context.Context, parser *gofeed.Parser, urlFeed *url.URL) (*gofeed.Feed, error) {
	resp, err := parser.ParseURLWithContext(urlFeed.String(), ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse rss feed: %w", err)
	}
	sort.Sort(resp)
	return resp, nil
}
