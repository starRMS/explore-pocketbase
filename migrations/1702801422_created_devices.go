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
			"id": "ycyfiura0pokvih",
			"created": "2023-12-17 08:23:42.292Z",
			"updated": "2023-12-17 08:23:42.292Z",
			"name": "devices",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "sn0voayy",
					"name": "name",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "bgqseth2",
					"name": "type",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "dslharui",
					"name": "capacity",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 1,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "te7njgay",
					"name": "verified",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "t2prbssb",
					"name": "organization",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "lvnkrzomlowy6g0",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "ozgcvxfw",
					"name": "longitude",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "f7eubqxd",
					"name": "latitude",
					"type": "number",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "feyurzfv",
					"name": "postal_code",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": 5,
						"max": 5,
						"pattern": "^[0-9]+$"
					}
				},
				{
					"system": false,
					"id": "ajhtwjer",
					"name": "country",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"ID",
							"SG"
						]
					}
				},
				{
					"system": false,
					"id": "rtkfww88",
					"name": "energy_source",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"solar",
							"wind",
							"hydro"
						]
					}
				},
				{
					"system": false,
					"id": "drzdyqbk",
					"name": "supported_by_goverment",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "pxmp9lvg",
					"name": "commission_date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "xze0kbfd",
					"name": "expiry_date",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "2000-01-01 12:00:00.000Z",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "pcjxdddw",
					"name": "issuer_notes",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [],
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

		collection, err := dao.FindCollectionByNameOrId("ycyfiura0pokvih")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
