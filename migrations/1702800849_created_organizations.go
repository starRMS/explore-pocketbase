package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "lvnkrzomlowy6g0",
			"created": "2023-12-17 08:14:09.138Z",
			"updated": "2023-12-17 08:14:09.138Z",
			"name": "organizations",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "nudungtp",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 5,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "aukxg8gz",
					"name": "address",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 10,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "eddxaarj",
					"name": "npwp",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 15,
						"max": 16,
						"pattern": "^[0-9]+"
					}
				},
				{
					"system": false,
					"id": "6kepsian",
					"name": "email",
					"type": "email",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"exceptDomains": [],
						"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "uzqlyc2i",
					"name": "phone_number",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 6,
						"max": 12,
						"pattern": "^021[0-9]{3,9}$"
					}
				},
				{
					"system": false,
					"id": "pujtyaqu",
					"name": "established_date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "1930-01-01 12:00:00.000Z",
						"max": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_unique_org_email` + "`" + ` ON ` + "`" + `organizations` + "`" + ` (` + "`" + `email` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_unique_org_npwp` + "`" + ` ON ` + "`" + `organizations` + "`" + ` (` + "`" + `npwp` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_unique_org_phone_number` + "`" + ` ON ` + "`" + `organizations` + "`" + ` (` + "`" + `phone_number` + "`" + `)"
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("lvnkrzomlowy6g0")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
