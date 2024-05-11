package swagger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func TableExists(db *sql.DB, table string) (bool, error) {

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name='%s') AS exists`, table)

	var exists bool
	err := db.QueryRow(query).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil

}

func RowWithIdExists(db *sql.DB, table string, id string) (bool, error) {

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s WHERE id=%s) AS exists`, table, id)

	var exists bool
	err := db.QueryRow(query).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil

}

func TableEmpty(db *sql.DB, table string) (bool, error) {

	query := fmt.Sprintf(`SELECT EXISTS (SELECT 1 FROM %s) AS exists`, table)

	var exists bool
	err := db.QueryRow(query).Scan(&exists)

	if err != nil {
		return false, err
	}

	return !exists, nil

}

func IdExists(db *sql.DB, table string, id string) (bool, error) {

	// Assert table exists //

	tableExists, err := TableExists(db, table)
	if err != nil {
		return false, err
	}

	if !tableExists {
		return false, nil
	}

	// Assert record exists //

	idExists, err := RowWithIdExists(db, table, id)
	if err != nil {
		log.Fatal(err)
	}

	if !idExists {
		return false, nil
	}

	return true, nil

}

func OpenDBConnection() *sql.DB {

	db, err := sql.Open("postgres", GetDbmsAddress())
	if err != nil {
		log.Fatal(err)
	}

	return db

}

type RowsToObject func(row *sql.Rows) map[string]interface{}

func GetAll(w http.ResponseWriter, r *http.Request, table string, rowToObject RowsToObject) {

	// Connect to DB //

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := OpenDBConnection()
	defer db.Close()

	// Assert table exists //

	exists, err := TableExists(db, table)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		w.Write([]byte("[]"))
		return
	}

	// Get all records //

	query := `SELECT * FROM ` + table

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Return records //

	var records []map[string]interface{}

	for rows.Next() {
		records = append(records, rowToObject(rows))
	}

	recordsJson, err := json.Marshal(records)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(recordsJson)

}

type AddToDB func(db *sql.DB) string

func Post(w http.ResponseWriter, r *http.Request, addRecord AddToDB) {

	// Connect to DB //

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := OpenDBConnection()
	defer db.Close()

	// Add to DB //

	responseBody := addRecord(db)

	// Respond //

	w.Write([]byte(responseBody))

}

type RowToObject func(row *sql.Row) map[string]interface{}

func GetById(w http.ResponseWriter, r *http.Request, table string, rowToObject RowToObject) {

	// Connect to DB //

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := OpenDBConnection()
	defer db.Close()

	// Get ID from path //

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid ID supplied", 400)
		return
	}

	// Assert ID exists //

	exists, err := IdExists(db, table, id)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		http.Error(w, "Item not found", 404)
		return
	}

	// Select record //

	row := db.QueryRow("SELECT * FROM " + table + " WHERE id=" + id)

	// Return record //

	jsonData, err := json.Marshal(rowToObject(row))
	if err != nil {
		log.Fatal(err)
	}

	w.Write(jsonData)

}

type PutInDB func(db *sql.DB, id string)

func PutById(w http.ResponseWriter, r *http.Request, table string, put PutInDB) {

	// Connect to DB //

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := OpenDBConnection()
	defer db.Close()

	// Get ID from path //

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid ID supplied", 400)
		return
	}

	// Assert ID exists //

	exists, err := IdExists(db, table, id)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		http.Error(w, "Item not found", 404)
		return
	}

	// Alter DB //

	put(db, id)

}

func DeleteById(w http.ResponseWriter, r *http.Request, table string) {

	// Connect to DB //

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	db := OpenDBConnection()
	defer db.Close()

	// Get ID from path //

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid ID supplied", 400)
		return
	}

	// Assert ID exists //

	exists, err := IdExists(db, table, id)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		http.Error(w, "Item not found", 404)
		return
	}

	// Delete record //

	query := fmt.Sprintf(
		`DELETE FROM %s WHERE id='%v'`, table, id)

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}
