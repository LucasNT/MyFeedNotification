package usecase

import (
	"context"
	"strings"

	"github.com/LucasNT/MyFeed/internal/adapters/mocks"
	"github.com/LucasNT/MyFeed/internal/entities"
)

func convertFeedToNotificationMessage(ctx context.Context, feed entities.Feed, appName string) (mocks.NotificationMessage, error) {
	var sb strings.Builder
	sb.WriteString("<a href=\"")
	sb.WriteString(feed.LinkToPage.String())
	sb.WriteString("\">")
	sb.WriteString(feed.Body)
	sb.WriteString("<\\a>")
	ret := mocks.NotificationMessage{
		Title:   feed.Title,
		AppName: appName,
		Level:   feed.Level,
		Summary: sb.String(),
	}
	return ret, nil

}
