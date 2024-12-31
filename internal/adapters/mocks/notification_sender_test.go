package mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
	"github.com/LucasNT/MyFeed/internal/entities"
	utilstest "github.com/LucasNT/MyFeed/internal/utils_test"
)

func TestNotificationSend(t *testing.T) {
	t.Run("Send a message, test if i get it", func(t *testing.T) {
		var want NotificationMessage = NotificationMessage{
			AppName: "My Feed Notification",
			Title:   "Ola mundo",
			Level:   entities.NORMAL_LEVEL,
			Summary: "Não sei como isso vai funcionar",
		}
		notificationSender := NewNotificationSend()
		err := notificationSender.Send(context.Background(), want.Level, want.Summary, want.Title, want.AppName)
		utilstest.AssertErrorIsNil(t, err)
		got := notificationSender.GetNotification()
		if want != got {
			t.Errorf("Test failed wanted %q, received %q", want, got)
		}
	})
	t.Run("Test error message", func(t *testing.T) {
		var n NotificationMessage = NotificationMessage{
			AppName: "My Feed Notification",
			Title:   "Ola mundo",
			Level:   entities.NORMAL_LEVEL,
			Summary: "Não sei como isso vai funcionar",
		}
		want := interfaces.ErrFailedToSendMessage
		notificationSender := NewNotificationSend()
		notificationSender.Fail = true
		got := notificationSender.Send(context.Background(), n.Level, n.Summary, n.Title, n.AppName)
		if !errors.Is(got, want) {
			t.Errorf("Wanted error %q, received %q", want, got)
		}
	})
	t.Run("Don't return the value in notification mock two times", func(t *testing.T) {
		var n NotificationMessage = NotificationMessage{
			AppName: "My Feed Notification",
			Title:   "Ola mundo",
			Level:   entities.NORMAL_LEVEL,
			Summary: "Não sei como isso vai funcionar",
		}
		want := NotificationMessage{}
		notificationSender := NewNotificationSend()
		err := notificationSender.Send(context.Background(), n.Level, n.Summary, n.Title, n.AppName)
		utilstest.AssertErrorIsNil(t, err)
		_ = notificationSender.GetNotification()
		got := notificationSender.GetNotification()

		if got != want {
			t.Errorf("Test Failed, expected empty valur got %q", got)
		}
	})
}
