/*
 * Inventory API
 *
 * This is a sample inventory management server based on the OpenAPI 3.0 specification
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {

	GetAll(w, r, "category", func(row *sql.Rows) map[string]interface{} {

		var id int
		var name string

		err := row.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		name = strings.TrimSpace(string(name))

		return map[string]interface{}{
			"id":   id,
			"name": name,
		}

	})

}

func PostCategory(w http.ResponseWriter, r *http.Request) {

	Post(w, r, func(db *sql.DB) string {

		// Parse request body //

		var categoryData CategoryData
		err := json.NewDecoder(r.Body).Decode(&categoryData)
		if err != nil {
			http.Error(w, "", 400)
			return ``
		}

		// Add record //

		var id int
		err = db.QueryRow(fmt.Sprintf(`INSERT INTO category (name) VALUES ('%v') RETURNING id`, categoryData.Name)).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		return fmt.Sprintf(`{"categoryId":%v}`, id)

	})

}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {

	GetById(w, r, "category", func(row *sql.Row) map[string]interface{} {

		var id int
		var name string

		err := row.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		name = strings.TrimSpace(string(name))

		return map[string]interface{}{
			"id":   id,
			"name": name,
		}

	})

}

func PutCategoryById(w http.ResponseWriter, r *http.Request) {

	PutById(w, r, "category", func(db *sql.DB, id string) {

		// Parse request body //

		var categoryData CategoryData

		err := json.NewDecoder(r.Body).Decode(&categoryData)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		// Update record //

		query := fmt.Sprintf(
			"UPDATE category SET name='%s' WHERE id=%v",
			categoryData.Name, id)

		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}

	})

}

func DeleteCategoryById(w http.ResponseWriter, r *http.Request) {

	DeleteById(w, r, "category")

}
