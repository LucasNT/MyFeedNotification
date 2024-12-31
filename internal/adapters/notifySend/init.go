package notifysend

import (
	"context"
	"fmt"
	"io"
	"os/exec"

	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
	"github.com/LucasNT/MyFeed/internal/entities"
)

type NotificationSend struct {
	command string
	configs []string
}

func New() *NotificationSend {
	ret := new(NotificationSend)
	ret.command = "notify-send"
	ret.configs = []string{"--app-name=My Rss Notification"}
	return ret
}

func convertLevelToNotifySendLevel(level int) string {
	var levelAux string
	switch level {
	case entities.LOW_LEVEL:
		levelAux = "low"
	case entities.NORMAL_LEVEL:
		levelAux = "normal"
	case entities.CRITICAL_LEVEL:
		levelAux = "critical"
	default:
		levelAux = "normal"
	}
	return levelAux
}

func (n *NotificationSend) Send(ctx context.Context, level int, body string, title string, appName string) error {
	args := make([]string, len(n.configs), len(n.configs)+4)
	copy(args, n.configs)
	var levelAux string = convertLevelToNotifySendLevel(level)
	args = append(args, "--urgency="+levelAux)
	args = append(args, "--app-name="+appName)
	args = append(args, title)
	args = append(args, body)
	command := exec.CommandContext(ctx, n.command, args...)
	ioErr, err := command.StderrPipe()
	if err != nil {
		return fmt.Errorf("%w, %w", interfaces.ErrFailedToSendMessage, err)
	}
	err = command.Start()
	if err != nil {
		return fmt.Errorf("%w, %w", interfaces.ErrFailedToSendMessage, err)
	}
	data, err := io.ReadAll(ioErr)
	if err != nil {
		return fmt.Errorf("%w, %w", interfaces.ErrFailedToSendMessage, err)
	}
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("%w, %w, %s", interfaces.ErrFailedToSendMessage, err, string(data))
	}
	return nil
}
