package main

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type dbSqLite struct {
	handler *sql.DB
}

/*
 id int
 alias string
 description string
 task_type string
 tags string
 timestamp int
 estimate_time string
 real_time string
 reminders string
*/

func (db dbSqLite) Create(t Task) error {

	sql_additem := `
    INSERT OR REPLACE INTO tasks(
    alias,
    description,
    task_type,
    tags,
    timestamp,
    estimate_time,
    real_time,
    reminders,
    InsertedDatetime
    ) values(?, ?, ?,?,?,?,?,?, CURRENT_TIMESTAMP)
    `

	stmt, err := db.handler.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(t.Alias, t.Description, t.Task_type, strings.Join(t.Tags, ","), t.Timestamp, t.Estimate_time, t.Real_time, strings.Join(t.Reminders, ","))
	if err2 != nil {
		panic(err2)
	}
	return nil

}

func (db *dbSqLite) ReadById(id *int64) (TaskList, error) {

	sql_readall := `
    SELECT id, alias, description, task_type, timestamp, estimate_time, real_time, tags, reminders FROM tasks
    WHERE id=?
    ORDER BY datetime(InsertedDatetime) DESC
    `
	stmt, err := db.handler.Prepare(sql_readall)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql_readall, id)

	//rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var result TaskList
	var dbTags, dbReminders string

	for rows.Next() {
		t := Task{}
		err2 := rows.Scan(&t.Id, &t.Alias, &t.Description, &t.Task_type, &t.Timestamp, &t.Estimate_time, &t.Real_time, &dbTags, &dbReminders)
		if err2 != nil {
			panic(err2)
		}
		t.Tags = strings.Split(dbTags, ",")
		t.Reminders = strings.Split(dbReminders, ",")
		result = append(result, t)
	}
	return result, err

}

func (db *dbSqLite) ReadByAlias(alias *string) (TaskList, error) {
	return nil, nil
}

func (db *dbSqLite) Update(t Task) error {
	return nil
}
func (db *dbSqLite) Delete(t Task) error {
	return nil
}

func (db *dbSqLite) ReadAll() (TaskList, error) {
	sql_readall := `
    SELECT id, alias, description, task_type, timestamp, estimate_time, real_time, tags, reminders FROM tasks
    ORDER BY datetime(InsertedDatetime) DESC
    `

	rows, err := db.handler.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result TaskList
	var dbTags, dbReminders string

	for rows.Next() {
		t := Task{}

		err2 := rows.Scan(&t.Id, &t.Alias, &t.Description, &t.Task_type, &t.Timestamp, &t.Estimate_time, &t.Real_time, &dbTags, &dbReminders)
		if err2 != nil {
			panic(err2)
		}
		t.Tags = strings.Split(dbTags, ",")
		t.Reminders = strings.Split(dbReminders, ",")
		result = append(result, t)
	}
	return result, err
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
    CREATE TABLE IF NOT EXISTS tasks(
    Id TEXT NOT NULL PRIMARY KEY,
    alias TEXT,
    description TEXT,
    tags Text,
    timestamp int,
    estimate_time DATETIME,
    real_time DATETIME,
    reminders TEXT,
    InsertedDatetime DATETIME
    );
    `

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}
