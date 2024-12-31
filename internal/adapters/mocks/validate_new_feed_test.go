package mocks

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/LucasNT/MyFeed/internal/entities"
	utilstest "github.com/LucasNT/MyFeed/internal/utils_test"
)

func TestValidate(t *testing.T) {
	t.Run("Validate data that not is in database", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		ok, err := validateNewFeed.Validate(context.Background(), feedOlamundo)
		utilstest.AssertErrorIsNil(t, err)
		if !ok {
			t.Errorf("Wanted for a \"true\" but received \"false\"")
		}
	})
	t.Run("Validate data that is in database, and old", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)
		addDuration, err := time.ParseDuration("1h")
		if err != nil {
			t.Fatalf("Failed to create duration for test, %q", err)
		}
		feedOlamundo.Time = feedOlamundo.Time.Add(addDuration)
		ok, err := validateNewFeed.Validate(context.Background(), feedOlamundo)
		if !ok {
			t.Errorf("Wanted for a \"true\" but received \"false\"")
		}
	})
	t.Run("Validate data that is in database, and is newer", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)
		addDuration, err := time.ParseDuration("-1h")
		if err != nil {
			t.Fatalf("Failed to create duration for test, %q", err)
		}
		feedOlamundo.Time = feedOlamundo.Time.Add(addDuration)
		ok, err := validateNewFeed.Validate(context.Background(), feedOlamundo)
		if ok {
			t.Errorf("Wanted for a \"true\" but received \"false\"")
		}
	})
	t.Run("Validate data that is in database, and there is more than one in the database", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)

		auxUrl, _ = url.Parse("https://status.cloud.google.com/")
		feedGCP := entities.NewFeed("Google cloud Status", "The cloud is on fire", time.Now(), auxUrl, entities.CRITICAL_LEVEL)
		validateNewFeed.WriteNewTime(context.Background(), feedGCP)

		addDuration, err := time.ParseDuration("-1h")
		if err != nil {
			t.Fatalf("Failed to create duration for test, %q", err)
		}
		feedOlamundo.Time = feedOlamundo.Time.Add(addDuration)
		ok, err := validateNewFeed.Validate(context.Background(), feedOlamundo)
		if ok {
			t.Errorf("Wanted for a \"true\" but received \"false\"")
		}
	})
	t.Run("Validate data that is in database, and is equal", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)
		ok, err := validateNewFeed.Validate(context.Background(), feedOlamundo)
		utilstest.AssertErrorIsNil(t, err)
		if ok {
			t.Errorf("Wanted for a \"true\" but received \"false\"")
		}
	})
}

func TestWriteNewTime(t *testing.T) {
	t.Run("Test if it write the data", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)
		err := validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)

		utilstest.AssertErrorIsNil(t, err)

		want := feedOlamundo
		got, ok := validateNewFeed.GetValue(want.Title)

		if !ok {
			t.Errorf("Failed to save data")
		} else if !got.Equal(want) {
			t.Errorf("Value saved %q is not equal to the value passed %q", got, want)
		}
	})
	t.Run("Test if it write multiple datas", func(t *testing.T) {
		validateNewFeed := NewValidateNewFeed()
		auxUrl, _ := url.Parse("https://ola.mundo.com/")
		feedOlamundo := entities.NewFeed("Ola mundo", "Como que vai funcionar", time.Now(), auxUrl, entities.NORMAL_LEVEL)

		auxUrl, _ = url.Parse("https://status.cloud.google.com/")
		feedGCP := entities.NewFeed("Google cloud Status", "The cloud is on fire", time.Now(), auxUrl, entities.CRITICAL_LEVEL)

		err := validateNewFeed.WriteNewTime(context.Background(), feedOlamundo)
		utilstest.AssertErrorIsNil(t, err)

		err = validateNewFeed.WriteNewTime(context.Background(), feedGCP)
		utilstest.AssertErrorIsNil(t, err)

		want := feedOlamundo
		got, ok := validateNewFeed.GetValue(want.Title)

		if !ok {
			t.Errorf("Failed to save data")
		} else if !got.Equal(want) {
			t.Errorf("Value saved %q is not equal to the value passed %q", got, want)
		}
	})
}
