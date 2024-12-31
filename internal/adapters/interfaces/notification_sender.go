package interfaces

import (
	"context"
	"errors"
)

var ErrFailedToSendMessage error = errors.New("Failed to send notification")

type NotificationSender interface {
	Send(ctx context.Context, level int, summary string, title string, appName string) error
}
