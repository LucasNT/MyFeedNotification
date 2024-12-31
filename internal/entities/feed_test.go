package entities

import (
	"net/url"
	"testing"
	"time"
)

func TestFeedEqual(t *testing.T) {
	t.Run("Test equal with same variable", func(t *testing.T) {
		link_path, _ := url.Parse("https://google.com.br/")
		feed := NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), link_path, NORMAL_LEVEL)

		if !feed.Equal(feed) {
			t.Errorf("Failed to compare feeds")
		}
	})
	t.Run("Test equal with two values equal", func(t *testing.T) {
		auxTime := time.Now()
		link_path, _ := url.Parse("https://google.com.br/")
		feed1 := NewFeed("Ola mundo", "Como que vai funcionar", auxTime, link_path, NORMAL_LEVEL)

		feed2 := NewFeed("Ola mundo", "Como que vai funcionar", auxTime, link_path, NORMAL_LEVEL)

		if !feed1.Equal(feed2) {
			t.Errorf("Expected that value %q and %q would be equal", feed1, feed2)
		}
	})

	t.Run("Test equal with two values not equal", func(t *testing.T) {
		auxTime := time.Now()
		link_path, _ := url.Parse("https://google.com.br/")
		feed1 := NewFeed("Ola mundo", "Como que vai funcionar", auxTime, link_path, NORMAL_LEVEL)

		feed2 := NewFeed("Hello Wolrd", "How this will worker", auxTime, link_path, NORMAL_LEVEL)

		if feed1.Equal(feed2) {
			t.Errorf("Expected that value %q and %q would not be equal", feed1, feed2)
		}
	})

	t.Run("Test equal with different levels", func(t *testing.T) {
		auxTime := time.Now()
		link_path, _ := url.Parse("https://google.com.br/")
		feed1 := NewFeed("Ola mundo", "Como que vai funcionar", auxTime, link_path, NORMAL_LEVEL)

		feed2 := NewFeed("Hello Wolrd", "How this will worker", auxTime, link_path, LOW_LEVEL)

		if feed1.Equal(feed2) {
			t.Errorf("Expected that value %q and %q would be equal", feed1, feed2)
		}
	})
}
