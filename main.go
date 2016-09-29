package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/abiosoft/river"
	_ "github.com/mattn/go-sqlite3"
    "gopkg.in/mgo.v2"
)


func main() {

	log.Println("Server init")

	db := &dbSqLite{handler: connect2Sqlite()}
    defer db.handler.Close()
    //db := &dbMongoDB{session: connect2Mongo(), dbName: "ws-0", collection:"tasks"}
    //defer db.session.Close()


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
func connect2Sqlite() *sql.DB {
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

func connect2Mongo() *mgo.Session {
    session, err := mgo.Dial("localhost, mongo.brcz.mk.ua")
    if err != nil {
        panic(err)
    }
    
    return session
}