package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "failed to execute command: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	go func() {
		select {
		case sig := <-sigCh:
			fmt.Fprintf(os.Stdout, "signal received: %v\n", sig)
			cancel()
		case <-ctx.Done():
		}
	}()

	for i := 1; i < 120; i++ {
		err := ctx.Err()
		if err != nil {
			fmt.Fprintf(os.Stdout, "context: %v\n", err)
			break
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
