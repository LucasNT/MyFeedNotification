package entities

import (
	"net/url"
	"testing"
	"time"
)

func TestFeedEqual(t *testing.T) {
	t.Run("Test equal with same variable", func(t *testing.T) {
		var feed Feed
		feed.Title = "Ola mundo"
		feed.Body = "Como que vai funcionar"
		feed.Time = time.Now()
		feed.LinkToPage, _ = url.Parse("https://google.com.br/")

		if !feed.Equal(feed) {
			t.Errorf("Failed to compare feeds")
		}
	})
	t.Run("Test equal with two values equal", func(t *testing.T) {
		auxTime := time.Now()
		var feed1 Feed
		feed1.Title = "Ola mundo"
		feed1.Body = "Como que vai funcionar"
		feed1.Time = auxTime
		feed1.LinkToPage, _ = url.Parse("https://google.com.br/")

		var feed2 Feed
		feed2.Title = "Ola mundo"
		feed2.Body = "Como que vai funcionar"
		feed2.Time = auxTime
		feed2.LinkToPage, _ = url.Parse("https://google.com.br/")

		if !feed1.Equal(feed2) {
			t.Errorf("Expected that value %q and %q would be equal", feed1, feed2)
		}
	})

	t.Run("Test equal with two values equal", func(t *testing.T) {
		auxTime := time.Now()
		var feed1 Feed
		feed1.Title = "Ola mundo"
		feed1.Body = "Como que vai funcionar"
		feed1.Time = auxTime
		feed1.LinkToPage, _ = url.Parse("https://google.com.br/")

		var feed2 Feed
		feed2.Title = "Hello Wolrd"
		feed2.Body = "How this will worker"
		feed2.Time = auxTime
		feed2.LinkToPage, _ = url.Parse("http://google.com.br/")

		if feed1.Equal(feed2) {
			t.Errorf("Expected that value %q and %q would not be equal", feed1, feed2)
		}
	})
}
