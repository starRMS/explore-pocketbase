package opentelemetry

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var shutdownFuncs []func(context.Context) error

func Init(ctx context.Context, jaegerURL string) (err error) {
	// Handles shutdown when error during
	// initializing such that no memory leaks
	handleErr := func(initErr error) {
		err = errors.Join(initErr, Shutdown(ctx))
	}

	// Resource
	resource, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("pocketbase"),
			semconv.ServiceVersion("v0.20.0"),
			semconv.ServiceInstanceID("pocketbase-app"),
			attribute.String("env", "development"),
		),
	)
	if err != nil {
		handleErr(err)
		return
	}

	// Trace Provider
	tp, err := newJaegerTracerProvider(jaegerURL, resource)
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
