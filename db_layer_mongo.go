package main

import (
        "fmt"
        //"log"
        "errors"
        
        "gopkg.in/mgo.v2"
        //"gopkg.in/mgo.v2/bson"
        )

type dbMongoDB struct {
    session *mgo.Session
    dbName string
    collection string
}

func (db *dbMongoDB) Create(t Task) error {
    c := db.session.DB(db.dbName).C(db.collection)
    err := c.Insert(&t)
    return err
}
func (db *dbMongoDB) ReadById(id *int64) (TaskList, error) {
    fmt.Println("mockDB.ReadById")
    return nil, errors.New("no such id")
}
func (db *dbMongoDB) ReadByAlias(alias *string) (TaskList, error) {
    fmt.Println("mockDB.ReadByAlias")
    task := mockTask(Task{Alias:"go-dms-workshop"}) //,Tags:["Golang", "Workshop", "DMS"]
    return TaskList{task}, nil
}
func (db *dbMongoDB) Update(t Task) error {
    return nil
}
func (db *dbMongoDB) Delete(t Task) error {
    return errors.New("Delete not supported")
}
func (db *dbMongoDB) ReadAll() (TaskList, error) {
    var tasks TaskList
    
    c := db.session.DB(db.dbName).C(db.collection)
    err := c.Find().All(&tasks)
    
    return tasks, err
}

