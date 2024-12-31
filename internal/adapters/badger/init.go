package badger

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LucasNT/MyFeed/internal/entities"
	db "github.com/dgraph-io/badger/v4"
)

type Badger struct {
	database *db.DB
}

func New(dbPath string) (Badger, error) {
	var err error
	ret := Badger{}
	ret.database, err = db.Open(db.DefaultOptions(dbPath))
	if err != nil {
		return Badger{}, fmt.Errorf("Failed to open database, %w", err)
	}
	return ret, nil
}

func (b Badger) Close() error {
	return b.database.Close()
}

func (b Badger) Validate(ctx context.Context, feed entities.Feed) (bool, error) {
	var buffer []byte = make([]byte, 64)
	err := b.database.View(func(txn *db.Txn) error {
		item, err := txn.Get([]byte(feed.Title))
		if err != nil {
			return err
		}
		buffer, err = item.ValueCopy(buffer)
		return err
	})
	if errors.Is(err, db.ErrKeyNotFound) {
		return true, nil
	} else if err != nil {
		return false, fmt.Errorf("Validate failed, failed to read database, %w", err)
	}
	var dateTime time.Time
	err = dateTime.UnmarshalBinary(buffer)
	if err != nil {
		return true, nil
	}

	return feed.Time.After(dateTime), nil
}

func (b Badger) WriteNewTime(ctx context.Context, feed entities.Feed) error {
	var value []byte
	var key []byte
	var err error
	value, err = feed.Time.MarshalBinary()
	key = []byte(feed.Title)
	err = b.database.Update(func(txn *db.Txn) error {
		return txn.Set(key, value)
	})
	if err != nil {
		return fmt.Errorf("Failed to write data to database %w", err)
	}
	return nil
}
