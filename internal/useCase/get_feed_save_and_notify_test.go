package usecase

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/LucasNT/MyFeed/internal/adapters/mocks"
	"github.com/LucasNT/MyFeed/internal/entities"
	utilstest "github.com/LucasNT/MyFeed/internal/utils_test"
)

func TestGetFeedSaveAndNotify(t *testing.T) {
	t.Run("Test if use case exeucte and don't return error", func(t *testing.T) {
		var mockGetFeed *mocks.GetFeed
		var mockValidateNewFeed *mocks.ValidateNewFeed
		var mockNotificationSender *mocks.NotificationSender

		_, wantUrlString, mockGetFeed, mockValidateNewFeed, mockNotificationSender := createUsecaseMocks()

		useCase := NewGetFeedSaveAndNotify(mockGetFeed, mockValidateNewFeed, mockNotificationSender)
		err := useCase.Execute(context.Background(), wantUrlString)
		utilstest.AssertErrorIsNil(t, err)
	})

	t.Run("Test if use case exeucte and send notification", func(t *testing.T) {
		var mockGetFeed *mocks.GetFeed
		var mockValidateNewFeed *mocks.ValidateNewFeed
		var mockNotificationSender *mocks.NotificationSender
		appName := "My Feed Notification Message Test"

		_, wantUrlString, mockGetFeed, mockValidateNewFeed, mockNotificationSender := createUsecaseMocks()

		useCase := NewGetFeedSaveAndNotify(mockGetFeed, mockValidateNewFeed, mockNotificationSender)
		useCase.AppName = appName
		err := useCase.Execute(context.Background(), wantUrlString)
		utilstest.AssertErrorIsNil(t, err)
		want := mocks.NotificationMessage{
			AppName: appName,
			Title:   "Ola mundo",
			Level:   entities.NORMAL_LEVEL,
			Summary: "<a href=\"https://ola.mundo.com/\">Como que vai funcionar<\\a>",
		}
		got := mockNotificationSender.GetNotification()
		if got != want {
			t.Errorf("Expected %q received %q", want, got)
		}

	})

	t.Run("Test if use case to not send old feeds", func(t *testing.T) {
		var mockGetFeed *mocks.GetFeed
		var mockValidateNewFeed *mocks.ValidateNewFeed
		var mockNotificationSender *mocks.NotificationSender
		appName := "My Feed Notification Message Test"

		feed, wantUrlString, mockGetFeed, mockValidateNewFeed, mockNotificationSender := createUsecaseMocks()

		useCase := NewGetFeedSaveAndNotify(mockGetFeed, mockValidateNewFeed, mockNotificationSender)
		useCase.AppName = appName
		err := useCase.Execute(context.Background(), wantUrlString)
		utilstest.AssertErrorIsNil(t, err)
		want := mocks.NotificationMessage{}
		_ = mockNotificationSender.GetNotification()
		feed.Time = feed.Time.Add(time.Duration(-1) * time.Hour)
		mockGetFeed.AppendFeed(feed, wantUrlString)

		err = useCase.Execute(context.Background(), wantUrlString)
		utilstest.AssertErrorIsNil(t, err)
		got := mockNotificationSender.GetNotification()
		if want != got {
			t.Errorf("Test Failed should returned a empty message returned %q", got)
		}
	})
}

func createUsecaseMocks() (
	feed entities.Feed,
	feedUrl string,
	getFeed *mocks.GetFeed,
	validateNewFeed *mocks.ValidateNewFeed,
	notificationSender *mocks.NotificationSender) {

	feedUrl = "https://ola.mundo.com/en/feed.atom"
	auxUrl, _ := url.Parse("https://ola.mundo.com/")
	feed = entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)

	getFeed = mocks.NewFeedGetter(feed, feedUrl)
	validateNewFeed = mocks.NewValidateNewFeed()
	notificationSender = mocks.NewNotificationSend()
	return
}
