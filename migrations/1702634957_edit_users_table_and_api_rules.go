package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// Up queries...
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		// API Rules
		collection.ListRule = nil
		collection.ViewRule = nil
		collection.CreateRule = nil
		collection.UpdateRule = nil
		collection.DeleteRule = nil

		collection.Indexes = types.JsonArray[string]{
			`CREATE UNIQUE INDEX idx_unique_user_nik ON users (nik)`,
		}

		collection.Schema.AddField(&schema.SchemaField{
			Name:        "nik",
			Type:        schema.FieldTypeText,
			Required:    true,
			Presentable: false,
			Unique:      true,
			Options: schema.TextOptions{
				Min:     types.Pointer(15),
				Max:     types.Pointer(16),
				Pattern: "^[0-9]+$",
			},
		})

		dao.SaveCollection(collection)

		return nil
	}, func(db dbx.Builder) error {
		// Down queries...
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		field := collection.Schema.GetFieldByName("nik")
		collection.Schema.RemoveField(field.Id)

		return nil
	})
}
