package opentelemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
)

var shutdownFuncs []func(context.Context) error

const (
	serviceName = "pocketbase"
	environment = "development"
	id          = 1
)

func Init(ctx context.Context, jaegerURL string) (err error) {
	// Handles shutdown when error during
	// initializing such that no memory leaks
	handleErr := func(initErr error) {
		err = errors.Join(initErr, Shutdown(ctx))
	}

	// Trace Provider
	tp, err := NewJaegerTracerProvider(jaegerURL)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tp.Shutdown)
	otel.SetTracerProvider(tp)

	return
}

func Shutdown(ctx context.Context) (err error) {
	for _, fn := range shutdownFuncs {
		if shutdownErr := fn(ctx); err != nil {
			err = errors.Join(err, shutdownErr)
		}
	}

	return err
}
