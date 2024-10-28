package useCase

import (
	"fmt"
	"io"
	"os"
	"time"
)

func (feed *Feed) TestIfFeedIsNew() (bool, error) {
	filePath, err := feed.convertUrlToPath()
	if err != nil {
		return false, fmt.Errorf("Failed to load feed state, %w", err)
	}
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("Failed to loadfeed state, %w", err)
	}
	defer file.Close()
	datas, err := io.ReadAll(file)
	if err != nil {
		return false, fmt.Errorf("Failed to loadfeed state, %w", err)
	}
	if len(datas) == 0 {
		return true, nil
	}
	var timeStamp time.Time
	timeStamp.UnmarshalBinary(datas)
	return feed.LastUpdate.After(timeStamp), nil
}
