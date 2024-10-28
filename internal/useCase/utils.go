package useCase

import (
	"os"
	"path"
	"strings"

	"github.com/adrg/xdg"
)

func (feed *Feed) convertUrlToPath() (string, error) {
	var dataFolder string = path.Join(xdg.DataHome, "myFeedNotificator")
	err := os.Mkdir(dataFolder, 0700)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	filePath := feed.FeedUrl.Hostname() + "/" + feed.FeedUrl.Path
	filePath = path.Clean(filePath)
	filePath = strings.ReplaceAll(filePath, "/", "%")
	filePath = strings.ReplaceAll(filePath, ":", "-")
	filePath = strings.ReplaceAll(filePath, ">", "-")
	filePath = strings.ReplaceAll(filePath, "<", "-")
	filePath = strings.ReplaceAll(filePath, "\"", "-")
	filePath = strings.ReplaceAll(filePath, "\\", "-")
	filePath = strings.ReplaceAll(filePath, "|", "-")
	filePath = strings.ReplaceAll(filePath, "?", "-")
	filePath = strings.ReplaceAll(filePath, "*", "-")
	filePath = path.Join(dataFolder, filePath)
	return filePath, nil
}
