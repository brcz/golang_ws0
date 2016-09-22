// Package main TODO API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// Provides API to operate with tasks
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /todo
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta


package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/abiosoft/river"
	_ "github.com/mattn/go-sqlite3"
)

//Step3: Implement of interaction with database
type dbDriver interface {
	Create(t Task) error
	ReadById(id *int64) (TaskList, error)
	ReadByAlias(alias *string) (TaskList, error)
	ReadAll() (TaskList, error)
	Update(t Task) error
	Delete(t Task) error
}

var db dbDriver

func main() {

	log.Println("Server init")

	db := &dbSqLite{handler: connect2Db()}

	rv := river.New()
	//Step2: Create API to handles such type of calls or use exists routes
	TODOHandler := river.NewEndpoint().
		Post("/", addTODORecord).
		Get("/", getTODOList).
		Get("/:id", getTODORecord).
		Put("/:id", updateTODORecord).
		Delete("/:id", deleteTODORecord)

	TODOHandler.Register(TODOModel(db))
	rv.Handle("/todo", TODOHandler)

	log.Println("Server ready. Listening on *:8081...")
	log.Fatal(http.ListenAndServe(":8081", rv))
}

//Step3: create connection with DB, docker-compose should be used for launch DB
func connect2Db() *sql.DB {
	db, err := sql.Open("sqlite3", "sq3_database.db")
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	CreateTable(db)
	return db
}
