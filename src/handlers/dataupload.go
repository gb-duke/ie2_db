package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gb-duke/ie2_db/src/dtos"
	"github.com/gorilla/mux"
)

type DataUploadStore struct {
	db *sql.DB
}

// Init initializes the DataStore
func (store *DataUploadStore) Init() {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"))

	// TODO: database type should be parameterized or an ENV var
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}

	store.db = db
}

// Get returns a user
func (store *DataUploadStore) Get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	du := dtos.DataUpload{}
	if len(id) == 0 {
		panic("404")
	}

	qry := fmt.Sprintf("SELECT du.id as id, du.userid as userid FROM datauploads WHERE du.id=%s", id)
	res, err := store.db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	res.Next()
	res.Scan(&du.ID, &du.UserID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(du)
}

// GetAll returns all the datauploads
func (store *DataUploadStore) GetAll(w http.ResponseWriter, r *http.Request) {

	var uploads []dtos.DataUpload

	sql := "SELECT * FROM datauploads"
	res, err := store.db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {

		var du dtos.DataUpload
		err := res.Scan(&du.ID, &du.UserID, &du.CreatedOn, &du.UpdatedOn, &du.DeletedOn, &du.UpdatedBy)

		if err != nil {
			panic(err.Error())
		}

		uploads = append(uploads, du)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploads)
}

// Create creates a new dataupload
func (store *DataUploadStore) Create(w http.ResponseWriter, r *http.Request) {
	var du dtos.DataUpload
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&du)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	curTime := time.Now().UTC().Format(time.RFC3339)
	du.CreatedOn = curTime
	du.UpdatedOn = curTime

	sql := fmt.Sprintf("INSERT INTO datauploads (userid, createdon, updatedon) VALUES ('%s', '%s', '%s')", du.UserID, du.CreatedOn, du.UpdatedOn)
	res, err := store.db.Query(sql)

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Update updates an existing dataupload
func (store *DataUploadStore) Update(w http.ResponseWriter, r *http.Request) {

	// TODO: implement
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

// Delete should delete a dataupload & possibly all related content
// I would recommend firing an RMQ message here to cascade deletes to interested listeners -gb
func (store *DataUploadStore) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	if len(id) == 0 {
		panic("404")
	}

	qry := fmt.Sprintf("DELETE * FROM datauploads WHERE datauploads.id=%s", id)
	res, err := store.db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	// TODO: add delete code
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
