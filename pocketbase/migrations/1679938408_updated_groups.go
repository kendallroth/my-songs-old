package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		// add
		new_type := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "kldn0vwb",
			"name": "type",
			"type": "select",
			"required": false,
			"unique": false,
			"options": {
				"maxSelect": 5,
				"values": [
					"A capella",
					"Instrumental",
					"Solo",
					"Orchestra",
					"Organ"
				]
			}
		}`), new_type)
		collection.Schema.AddField(new_type)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("m617x3t48r2h4bk")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("kldn0vwb")

		return dao.SaveCollection(collection)
	})
}
