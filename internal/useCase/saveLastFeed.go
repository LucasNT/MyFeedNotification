package useCase

import (
	"context"
	"fmt"
	"os"
)

func (feed *Feed) SaveLastFeed(ctx context.Context) (bool, error) {
	filePath, err := feed.convertUrlToPath()
	if err != nil {
		return false, fmt.Errorf("Failed to save feed state, %w", err)
	}
	binaryTime, err := feed.LastUpdate.MarshalBinary()
	if err != nil {
		return false, fmt.Errorf("Failed to save feed state, %w", err)
	}
	err = os.WriteFile(filePath, binaryTime, 0600)
	if err != nil {
		return false, fmt.Errorf("Failed to save feed state, %w", err)
	}
	return true, nil
}
