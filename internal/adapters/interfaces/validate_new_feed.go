package interfaces

import (
	"context"

	"github.com/LucasNT/MyFeed/internal/entities"
)

type ValidateNewFeed interface {
	Validate(ctx context.Context, feed entities.Feed) (bool, error)
	WriteNewTime(ctx context.Context, feed entities.Feed) error
}
