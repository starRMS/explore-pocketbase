package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/starRMS/explore-pocketbase/hooks"
	"github.com/starRMS/explore-pocketbase/pkg/observability"
	"github.com/starRMS/explore-pocketbase/routers/middlewares"
	"github.com/starRMS/explore-pocketbase/tools/writer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	// Import migrations
	_ "github.com/starRMS/explore-pocketbase/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	app := pocketbase.New()
	writer := writer.NewWriter()

	/*
		*************************************************
		Bootstrap Server - Initializing packages
		*************************************************
	*/
	app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
		if err := observability.Init(ctx); err != nil {
			log.Fatalf("error when initializing opentelemetry %s\n", err)
		}

		return nil
	})

	/*
		*************************************************
		Before Serve - Adding routes/middlewares
		*************************************************
	*/
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		meter := otel.Meter("http.server")
		histogram, err := meter.Float64Histogram("http.server.response.latency",
			metric.WithDescription("HTTP response latency"),
			metric.WithUnit("ms"),
		)
		if err != nil {
			writer.Logf("error initializing http latency meter: %s\n", err)
			return err
		}

		// Metric middleware to record latency
		e.Router.Use(middlewares.MetricMiddleware(histogram))
		return nil
	})

	/*
		*************************************************
		App Terminate - Handle shutdown properly
		*************************************************
	*/
	app.OnTerminate().Add(func(e *core.TerminateEvent) error {
		shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()
		// Gracefully shutdown opentelemetry
		writer.Log("Shutting down opentelemetry...")
		if err := observability.Shutdown(shutdownCtx); err != nil {
			writer.Errorf("error while shutting down opentelemetry %s\n", err)
			return err
		}
		return nil
	})

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// Auto created migration files from the
		// admin UI only auto migrate if go run
		Automigrate: strings.HasPrefix(os.Args[0], os.TempDir()),
	})

	/*
		*************************************************
		Router Static Files
		*************************************************
	*/
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	/*
		*************************************************
		User Hooks
		*************************************************
	*/
	app.OnModelAfterCreate("users").Add(hooks.User.ModelAfterCreate(writer, otel.Tracer("user-trace")))

	/*
		*************************************************
		Start the app
		*************************************************
	*/
	// FIXME: No need graceful shutdown pocketbase app??
	pbServeError := make(chan error, 1)
	go func() {
		pbServeError <- app.Start()
	}()

	select {
	case <-pbServeError:
		return
	case <-ctx.Done():
		writer.Log("Shutting down application...")
		stop()
	}
}
