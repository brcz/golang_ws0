package main

import "strconv"

// model file can interact with REST service and with DB

type Task struct {
	Id            int64    `json:"id,omitempty"`
	Alias         string   `json:"alias,omitempty"`
	Description   string   `json:"desc,omitempty"`
	Task_type     string   `json:"type,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Timestamp     int32    `json:"ts,omitempty"`
	Estimate_time string   `json:"etime,omitempty"`
	Real_time     string   `json:"rtime,omitempty"`
	Reminders     []string `json:"reminders,omitempty"`
}

type TaskList []Task

var Tasks TaskList

type Model struct {
	get    func(id string) interface{}
	getAll func() interface{}
	add    func(items ...interface{})
	put    func(id string, item interface{})
	delete func(id string)
}

func TODOModel(db dbDriver) Model {

	var model Model

	model.add = func(items ...interface{}) {
		for i := range items {
			//Tasks = append(Tasks, items[i].(Task))
			db.Create(items[i].(Task))
		}
	}

	model.getAll = func() interface{} {
		tasks, err := db.ReadAll()
		if err != nil {
			return nil
		}
		return tasks
	}

	model.get = func(id string) interface{} {
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil
		}

		if tasks, err := db.ReadById(&int64_id); err == nil {
			return tasks
		}

		return nil
	}

	return model
}
