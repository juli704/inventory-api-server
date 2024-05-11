package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	sw "inventory-api-server/go"
	"log"
	"net/http"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func clearTable(t *testing.T, table string) {

	// Request all records //

	resp, err := http.Get(fmt.Sprintf(`%s/%s`, sw.GetApiServerAddress(), table))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// Decode response body //

	type Id struct {
		Id int `json:"id"`
	}

	var objects []Id
	err = json.NewDecoder(resp.Body).Decode(&objects)
	if err != nil {
		t.Error(err)
	}

	// Erase all records //

	for _, obj := range objects {

		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(`%s/%s/%v`, sw.GetApiServerAddress(), table, obj.Id), nil)
		if err != nil {
			t.Error(err)
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Error("Clear failed with status code", resp.StatusCode)
		}

	}

	// Reset primary key serial counter //

	db, err := sql.Open("postgres", sw.GetDbmsAddress())
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	_, err = db.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1")
	if err != nil {
		t.Error(err)
	}

}

func PopulateTable(t *testing.T, table string, postBodies []string) {

	for _, body := range postBodies {

		// Post //

		requestBody := bytes.NewBuffer([]byte(body))

		resp, err := http.Post(fmt.Sprintf(`%s/%s`, sw.GetApiServerAddress(), table), "application/json", requestBody)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Error("Populate failed with status code", resp.StatusCode)
		}

	}

}

type TableSetup struct {
	table      string
	postBodies []string
}

func SetupTest(t *testing.T, tableSetups []TableSetup) func(t *testing.T) {

	// Load ENV file
	if err := godotenv.Load("../default.env"); err != nil {
		log.Print("No .env file found")
	}

	// Clear all tables that were set up.
	// In reverse because of foreign key constraints
	for i := len(tableSetups) - 1; i >= 0; i-- {
		clearTable(t, tableSetups[i].table)
	}

	// Populate all tables
	for _, setup := range tableSetups {
		PopulateTable(t, setup.table, setup.postBodies)
	}

	// Return tear down procedure
	return func(t *testing.T) {

		// Clear all tables that were set up.
		// In reverse because of foreign key constraints
		for i := len(tableSetups) - 1; i >= 0; i-- {
			clearTable(t, tableSetups[i].table)
		}

	}

}
