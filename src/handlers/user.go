package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gb-duke/ie2_db/src/dtos"
	"github.com/gorilla/mux"
)

// User should store and retrieve user data
type UsersStore struct {
	db *sql.DB
}

// Init initializes the DataStore
func (store *UsersStore) Init() {

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
func (store *UsersStore) Get(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	user := dtos.User{}
	if len(id) == 0 {
		panic("404")
	}

	qry := fmt.Sprintf("SELECT users.id as id, users.firstname as firstname, users.lastname as lastname, users.email as email FROM users WHERE users.id='%s'", id)
	res, err := store.db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	res.Next()
	res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAll returns all the users
func (store *UsersStore) GetAll(w http.ResponseWriter, r *http.Request) {

	qry := "SELECT users.id, users.firstname, users.lastname, users.email FROM users"
	res, err := store.db.Query(qry)
	var users []dtos.User
	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	for res.Next() {

		var user dtos.User
		err := res.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Create creates a new user
func (store *UsersStore) Create(w http.ResponseWriter, r *http.Request) {
	var user dtos.User
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update updates an existing user
func (store *UsersStore) Update(w http.ResponseWriter, r *http.Request) {

	// TODO: add update code
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}

// Delete should delete a user & possibly all related content
// I would recommend firing an RMQ message here to cascade deletes to interested listeners -gb
func (store *UsersStore) Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	if len(id) == 0 {
		panic("404")
	}

	qry := fmt.Sprintf("DELETE * FROM users WHERE users.id=%s", id)
	res, err := store.db.Query(qry)

	if err != nil {
		panic(err.Error())
	}

	defer res.Close()

	// TODO: add delete code
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
