package mocks

import (
	"context"

	"github.com/LucasNT/MyFeed/internal/entities"
)

type ValidateNewFeed struct {
	Value map[string]entities.Feed
}

func NewValidateNewFeed() *ValidateNewFeed {
	ret := new(ValidateNewFeed)
	ret.Value = make(map[string]entities.Feed)
	return ret
}

func (v *ValidateNewFeed) Validate(ctx context.Context, feed entities.Feed) (bool, error) {
	value, ok := v.Value[feed.Title]
	if ok {
		return value.Time.Before(feed.Time), nil
	}
	return true, nil
}

func (v *ValidateNewFeed) WriteNewTime(ctx context.Context, feed entities.Feed) error {
	v.Value[feed.Title] = feed
	return nil
}

func (v *ValidateNewFeed) GetValue(key string) (entities.Feed, bool) {
	ret, ok := v.Value[key]
	return ret, ok
}
