package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oking02/signapse-test/internal/application"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	a, err := application.Setup()
	if err != nil {
		panic(err)
	}

	go func() {
		<-ctx.Done()
		<-time.After(5 * time.Second)
	}()

	if err = a.Run(ctx); err != nil {
		panic(err)
	}

	os.Exit(1)
}
