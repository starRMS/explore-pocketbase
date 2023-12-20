package observability

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

var shutdownFuncs []func(context.Context) error

func Init(ctx context.Context) (err error) {
	// Handles shutdown when error during
	// initializing such that no memory leaks
	handleErr := func(initErr error) {
		err = errors.Join(initErr, Shutdown(ctx))
	}

	// Resource -.
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

	// Trace Exporter -.
	traceExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		handleErr(err)
		return
	}

	// Trace Provider -.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(resource),
	)
	shutdownFuncs = append(shutdownFuncs, tp.Shutdown)
	otel.SetTracerProvider(tp)

	// Metrics Exporter -.
	metricsExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint("localhost:4318"),
		otlpmetrichttp.WithInsecure(),
	)
	if err != nil {
		handleErr(err)
		return
	}

	// Meter Provider -.
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(metricsExporter, sdkmetric.WithInterval(15*time.Second)),
		),
	)
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetMeterProvider(mp)

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
