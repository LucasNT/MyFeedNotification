package useCase

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

// TODO refactor this to be better
func (feed *Feed) Notify(ctx context.Context) error {
	s := fmt.Sprintf(`<a href="%s"> %s </a>`, feed.Link, feed.Message)
	command := exec.CommandContext(ctx, "notify-send", "--app-name=My RSS Notification", feed.Title, s)
	ioErr, err := command.StderrPipe()
	if err != nil {
		return fmt.Errorf("Failed notification %w", err)
	}
	err = command.Start()
	if err != nil {
		return fmt.Errorf("Failed notification %w", err)
	}
	data, err := io.ReadAll(ioErr)
	if err != nil {
		return fmt.Errorf("Failed notification %w", err)
	}
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("Failed notification %w, %s", err, string(data))
	}
	return err
}
