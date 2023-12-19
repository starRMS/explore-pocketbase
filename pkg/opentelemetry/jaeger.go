package opentelemetry

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func newJaegerTracerProvider(url string, resource *resource.Resource) (*trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	), nil
}
