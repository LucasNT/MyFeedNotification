package useCase

import (
	"net/url"
	"time"
)

type DatabaseConnection struct {
	//TODO do this thing
}

type Feed struct {
	FeedUrl    *url.URL
	LastUpdate *time.Time
	Message    string
	Title      string
	Link       string
}
