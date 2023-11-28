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

	if err := srv.Run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)

		srv.Close()
		os.Exit(1)
	}

	<-ctx.Done()

	fmt.Fprintln(os.Stderr, "context done")
	srv.Close()
}
