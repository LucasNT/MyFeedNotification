package interfaces

import (
	"context"
	"fmt"
	"net/url"

	"github.com/LucasNT/MyFeed/internal/entities"
)

type FeedGetter interface {
	GetFeed(ctx context.Context, link *url.URL) (entities.Feed, error)
}

var (
	ErrNofeedToReturn error = fmt.Errorf("Didn't find the Feed with the link provided")
)
