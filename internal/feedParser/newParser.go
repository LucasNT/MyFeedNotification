package feedParser

import "github.com/mmcdole/gofeed"

func NewParser() *gofeed.Parser {
	return gofeed.NewParser()
}
