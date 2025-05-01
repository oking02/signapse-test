package application

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Component interface {
	Run(ctx context.Context) error
}

type App struct {
	components []Component
}

func (a *App) Run(ctx context.Context) error {

	eg, egCtx := errgroup.WithContext(ctx)

	for _, c := range a.components {
		component := c
		eg.Go(func() error {
			return component.Run(egCtx)
		})
	}
	return eg.Wait()
}
