package main

import (
	"context"
	"net/url"
	"sync"
	"time"

	"github.com/LucasNT/MyFeed/internal/useCase"
	"github.com/sirupsen/logrus"
)

func main() {
	var ctx context.Context = context.Background()
	var ctxTimeout context.Context
	var cancelFunction context.CancelFunc
	var wg sync.WaitGroup
	ctxTimeout, cancelFunction = context.WithTimeout(ctx, 10*time.Second)
	defer cancelFunction()

	urlFeed, err := url.Parse("https://status.cloud.google.com/en/feed.atom")
	if err != nil {
		logrus.Fatal(err)
	}
	wg.Add(1)

	go execute(ctxTimeout, urlFeed, &wg)

	wg.Wait()
}

func execute(ctx context.Context, urlFeed *url.URL, wg *sync.WaitGroup) {
	defer wg.Done()

	feed := useCase.NewFeed(urlFeed)

	err := feed.GetFeed(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}

	test, err := feed.TestIfFeedIsNew()
	if err != nil {
		logrus.Error(err)
		return
	}

	if test {
		_, err = feed.SaveLastFeed(ctx)
		if err != nil {
			logrus.Error(err)
			return
		}

		err = feed.Notify(ctx)
		if err != nil {
			logrus.Error(err)
			return
		}
	}
}
