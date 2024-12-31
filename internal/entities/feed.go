package entities

import (
	"net/url"
	"time"
)

type Feed struct {
	Title      string
	Body       string
	Time       time.Time
	LinkToPage *url.URL
	Level      int
}

func NewFeed(title string, body string, time time.Time, linkToPage *url.URL, level int) Feed {
	var ret Feed
	ret.Title = title
	ret.Body = body
	ret.Time = time
	ret.LinkToPage = linkToPage
	ret.Level = level
	return ret
}

func (f Feed) Equal(g Feed) bool {
	return f.Title == g.Title &&
		f.Body == g.Body &&
		f.Time == g.Time &&
		f.LinkToPage.Hostname() == g.LinkToPage.Hostname() &&
		f.LinkToPage.EscapedPath() == g.LinkToPage.EscapedPath()
}
