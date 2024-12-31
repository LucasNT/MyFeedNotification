package usecase

import (
	"context"
	"fmt"
	"net/url"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
)

type GetFeedSaveAndNotify struct {
	getFeed            interfaces.FeedGetter
	validateNewFeed    interfaces.ValidateNewFeed
	notificationSender interfaces.NotificationSender
	AppName            string
}

func NewGetFeedSaveAndNotify(getFeed interfaces.FeedGetter,
	ValidateNewFeed interfaces.ValidateNewFeed,
	NotificationSender interfaces.NotificationSender) GetFeedSaveAndNotify {

	var ret GetFeedSaveAndNotify
	ret.getFeed = getFeed
	ret.validateNewFeed = ValidateNewFeed
	ret.notificationSender = NotificationSender

	return ret
}

func (g GetFeedSaveAndNotify) Execute(ctx context.Context, link string) error {
	feedCtx, feedCtxCancelFunc := context.WithCancel(ctx)
	defer feedCtxCancelFunc()
	url, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("Failed to execute use case, invalid link: %w", err)
	}
	feed, err := g.getFeed.GetFeed(feedCtx, url)
	if err != nil {
		return fmt.Errorf("Failed to execute use case, parser Error: %w", err)
	}

	validateCtx, validateCtxCancelFunc := context.WithCancel(ctx)
	defer validateCtxCancelFunc()
	ok, err := g.validateNewFeed.Validate(validateCtx, feed)

	if err != nil {
		return fmt.Errorf("Failed to execute use case, validation time failed %w", err)
	}

	if !ok {
		return nil
	}

	err = g.validateNewFeed.WriteNewTime(validateCtx, feed)
	if err != nil {
		return fmt.Errorf("Failed to execute use case, write new time %w", err)
	}

	message, err := convertFeedToNotificationMessage(ctx, feed, g.AppName)

	notificationCtx, notificationCtxCancelFunc := context.WithCancel(ctx)
	defer notificationCtxCancelFunc()
	err = g.notificationSender.Send(notificationCtx, message.Level, message.Summary, message.Title, message.AppName)
	if err != nil {
		return fmt.Errorf("Failed to execute use case, notification error %w", err)
	}

	return nil
}
