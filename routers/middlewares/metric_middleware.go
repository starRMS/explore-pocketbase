package middlewares

import (
	"time"

	"github.com/labstack/echo/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func MetricMiddleware(histogram metric.Float64Histogram) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			defer func() {
				duration := float64(time.Since(start).Milliseconds())

				histogram.Record(c.Request().Context(), duration,
					metric.WithAttributes(
						attribute.String("http.request.path", c.Path()),
						attribute.String("http.request.method", c.Request().Method),
					),
				)
			}()

			return next(c)
		}
	}
}
