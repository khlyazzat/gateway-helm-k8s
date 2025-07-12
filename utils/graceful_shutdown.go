package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulShutdown(ctx context.Context) (context.Context, context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	return context.WithTimeout(ctx, 10*time.Second)
}
