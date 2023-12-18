package main

import (
	"log"
	"os"
	"strings"

	"github.com/starRMS/explore-pocketbase/hooks"
	"github.com/starRMS/explore-pocketbase/tools/writer"

	// Import migrations
	_ "github.com/starRMS/explore-pocketbase/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	writer := writer.NewWriter()

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
	app.OnModelAfterCreate("users").Add(hooks.User.ModelAfterCreate(writer))

	/*
		*************************************************
		Start the app
		*************************************************
	*/
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
