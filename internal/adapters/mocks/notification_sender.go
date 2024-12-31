package mocks

import (
	"context"
	"fmt"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
)

type NotificationMessage struct {
	AppName string
	Title   string
	Level   int
	Summary string
}

type NotificationSender struct {
	buffer NotificationMessage
	Fail   bool
}

func NewNotificationSend() *NotificationSender {
	ret := new(NotificationSender)
	ret.Fail = false
	return ret
}

func (n *NotificationSender) Send(ctx context.Context, level int, summary string, title string, appName string) error {
	if n.Fail {
		return fmt.Errorf("%w", interfaces.ErrFailedToSendMessage)
	} else {
		n.buffer = NotificationMessage{
			Level:   level,
			Title:   title,
			Summary: summary,
			AppName: appName,
		}
	}
	return nil
}

func (n *NotificationSender) GetNotification() NotificationMessage {
	ret := n.buffer
	n.buffer = NotificationMessage{}
	return ret
}
