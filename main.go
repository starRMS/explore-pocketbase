package main

import (
	"log"
	"os"
	"strings"

	"github.com/starRMS/explore-pocketbase/tools/encryptor"
	"github.com/starRMS/explore-pocketbase/tools/writer"

	// Import migrations
	_ "github.com/starRMS/explore-pocketbase/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
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
		On Terminate Hooks
		*************************************************
	*/
	app.OnTerminate().PreAdd(func(e *core.TerminateEvent) error {
		// Add logs for terminating
		println()
		writer.Log("Terminating...\n")
		writer.Log("Thanks for using pocketbase...\n")
		return nil
	})

	/*
		*************************************************
		User Hooks
		*************************************************
	*/
	app.OnModelAfterCreate("users").Add(func(e *core.ModelEvent) error {
		// Changes the value of user NIK after created.
		id := e.Model.GetId()

		return e.Dao.RunInTransaction(func(tx *daos.Dao) error {
			record, err := tx.FindRecordById("users", id)
			if err != nil {
				writer.Error("OnModelAfterCreate - FindRecordById: %s\n", err)
				return err
			}

			nik := record.Get("nik").(string)
			encrypted, err := encryptor.AES_CBC_Encrypt(nik)
			if err != nil {
				writer.Error("unable to encrypt NIK: %s\n", err)
				return err
			}
			record.Set("nik", encrypted)

			if err := tx.SaveRecord(record); err != nil {
				writer.Error("OnModelAfterCreate - SaveRecord: %s\n", err)
				return err
			}

			return nil

			// -- Another way to do it --
			//
			// user := domain.User{}
			//
			// if err := tx.DB().NewQuery(`SELECT id, nik FROM users WHERE id = {:id}`).
			// 	Bind(dbx.Params{
			// 		"id": id,
			// 	}).One(&user); err != nil {
			// 	writer.Error("UserOnModelAfterCreate %s\n", err)
			// 	return err
			// }
			//
			// user.NIK = "NIK_modified_using_custom_pocketbase_hooks"
			//
			// if _, err := tx.DB().NewQuery(`UPDATE users SET nik = {:nik} WHERE id = {:id}`).Bind(dbx.Params{
			// 	"nik": user.NIK,
			// 	"id":  user.Id,
			// }).Execute(); err != nil {
			// 	writer.Error("UserOnModelAfterCreate %s\n", err)
			// 	return err
			// }
			//
			// return nil
		})
	})

	/*
		*************************************************
		Start the app
		*************************************************
	*/
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
