package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		sig := <-sigs

		_, _ = fmt.Fprintf(os.Stderr, "Shutting down server. Reason: %s...\n", sig.String())

		cancel()
	}()

	srv := NewServer()

	errChan := make(chan error, 1)
	errChan <- srv.Run(ctx)

	select {
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, "context done")

		srv.Close(ctx)
	case err := <-errChan:
		fmt.Fprintln(os.Stderr, err)

		srv.Close(ctx)
		os.Exit(1)
	}
}
