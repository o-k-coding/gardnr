package cmdutil

import (
	"context"
	"os"
	"os/signal"
)

func ContextWithSignal(sigs ...os.Signal) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sigs...)
	go func() {
		<-ch
		cancel()
	}()

	return ctx, cancel
}
