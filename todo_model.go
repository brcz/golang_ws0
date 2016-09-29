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

//type TaskTags  string

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
			if tasks, err2 := db.ReadByAlias(&id); err2 == nil {
				return tasks
			}
			return err
		}

		if tasks, err := db.ReadById(&int64_id); err == nil {
			return tasks
		}

		return err
	}

	model.put = func(id string, item interface{}) {
		var updateTask Task
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		updateTask  = item.(Task)
		updateTask.Id = int64_id
		err = db.Update(updateTask)
		if err != nil {
			return
		}
		return
	}

	model.delete = func(id string) {
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return
		}
		err = db.Delete(Task{Id:int64_id})
		if err != nil {
			return
		}
		return
	}

	return model
}
