package mocks

import (
	"context"
	"errors"
	"net/url"
	"testing"
	"time"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
	"github.com/LucasNT/MyFeed/internal/entities"
	utilstest "github.com/LucasNT/MyFeed/internal/utils_test"
)

func TestGetFeed(t *testing.T) {
	t.Run("Test if one element", func(t *testing.T) {
		auxUrl, _ := url.Parse("https://google.com.br/")
		linkGet, _ := url.Parse("https://ola.mundo.com/")
		want := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		feed := NewFeedGetter(want, "https://ola.mundo.com/")
		got, err := feed.GetFeed(context.Background(), linkGet)

		utilstest.AssertErrorIsNil(t, err)
		assertFeed(t, got, want)
	})
	t.Run("Test get one of two elements", func(t *testing.T) {
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		auxUrl, _ = url.Parse("https://status.cloud.google.com/")
		feedGCP := entities.NewFeed("Google cloud Status", "The cloud is on fire", time.Now(), auxUrl, entities.CRITICAL_LEVEL)

		urlOlaMundo := "https://ola.mundo.com/en/feed.atom"
		feedGeeter := NewFeedGetter(feedOlamundo, urlOlaMundo)
		urlGCP := "https://status.cloud.google.com/en/feed.atom"
		feedGeeter.AppendFeed(feedGCP, urlGCP)

		want := feedGCP
		urlFeed, _ := url.Parse(urlGCP)
		got, err := feedGeeter.GetFeed(context.Background(), urlFeed)
		utilstest.AssertErrorIsNil(t, err)
		assertFeed(t, got, want)
	})
	t.Run("Test get of two elements", func(t *testing.T) {
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		auxUrl, _ = url.Parse("https://status.cloud.google.com/")
		feedGCP := entities.NewFeed("Google cloud Status", "The cloud is on fire", time.Now(), auxUrl, entities.CRITICAL_LEVEL)

		urlOlaMundo := "https://ola.mundo.com/en/feed.atom"
		feedGeeter := NewFeedGetter(feedOlamundo, urlOlaMundo)
		urlGCP := "https://status.cloud.google.com/en/feed.atom"
		feedGeeter.AppendFeed(feedGCP, urlGCP)

		want := feedGCP
		urlFeed, _ := url.Parse(urlGCP)
		got, err := feedGeeter.GetFeed(context.Background(), urlFeed)
		utilstest.AssertErrorIsNil(t, err)
		assertFeed(t, got, want)

		want = feedOlamundo
		urlFeed, _ = url.Parse(urlOlaMundo)
		got, err = feedGeeter.GetFeed(context.Background(), urlFeed)
		utilstest.AssertErrorIsNil(t, err)
		assertFeed(t, got, want)
	})
	t.Run("Test with a link without feed", func(t *testing.T) {
		urlString := "https://ola.mundo.com/"
		linkGet, _ := url.Parse(urlString)
		want := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), linkGet, entities.NORMAL_LEVEL)
		feed := NewFeedGetter(want, "https://status.cloud.google.com/en/feed.atom")
		_, err := feed.GetFeed(context.Background(), linkGet)

		if !errors.Is(err, interfaces.ErrNofeedToReturn) {
			t.Errorf("Expected wrapped error if %q, received %q", interfaces.ErrNofeedToReturn, err)
		}
	})
}

func assertFeed(t testing.TB, got, want entities.Feed) {
	t.Helper()
	if !got.Equal(want) {
		t.Errorf("got %q wanted %q", got, want)
		t.FailNow()
	}
}
