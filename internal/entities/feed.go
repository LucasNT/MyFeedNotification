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
}

func NewFeed(title string, body string, time time.Time, linkToPage *url.URL) *Feed {
	ret := new(Feed)
	ret.Title = title
	ret.Body = body
	ret.Time = time
	ret.LinkToPage = linkToPage
	return ret
}

func (f Feed) Equal(g Feed) bool {
	return f.Title == g.Title &&
		f.Body == g.Body &&
		f.Time == g.Time &&
		f.LinkToPage.Hostname() == g.LinkToPage.Hostname() &&
		f.LinkToPage.EscapedPath() == g.LinkToPage.EscapedPath()
}
