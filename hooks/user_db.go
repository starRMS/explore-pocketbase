package hooks

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/starRMS/explore-pocketbase/tools/encryptor"
	"github.com/starRMS/explore-pocketbase/tools/writer"
)

type userHooks int

var User userHooks

func (userHooks) ModelAfterCreate(writer *writer.Writer) func(e *core.ModelEvent) error {
	return func(e *core.ModelEvent) error {
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
	}
}
