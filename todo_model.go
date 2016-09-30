package main

import "strconv"

// model file can interact with REST service and with DB
// swagger:model
type Task struct {
	// the id for this user
	//
	// required: true
	// min: 1
	Id int64 `required json:"id,omitempty"`

	// required: true
	// min length: 3
	Alias string `json:"alias,omitempty"`

	// required: true
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

type modelResult struct {
	result interface{}
	status error
}

type Model struct {
	get    func(id string) modelResult
	getAll func() modelResult
	add    func(items ...interface{}) modelResult
	put    func(id string, item interface{}) modelResult
	delete func(id string) modelResult
}

func TODOModel(db dbDriver) Model {

	var model Model

	model.add = func(items ...interface{}) modelResult {
		for i := range items {
			err := db.Create(items[i].(Task))
			if err != nil {
				return modelResult{nil, err}
			}
		}
		return modelResult{items, nil}
	}

	model.getAll = func() modelResult {
		tasks, err := db.ReadAll()
		if err != nil {
			return modelResult{nil, err}
		}
		return modelResult{tasks, nil}
	}

	model.get = func(id string) modelResult {
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			if tasks, err2 := db.ReadByAlias(&id); err2 == nil {
				return modelResult{tasks, nil}
			}
			return modelResult{nil, err}
		}

		if tasks, err := db.ReadById(&int64_id); err == nil {
			return modelResult{tasks, nil}
		}

		return modelResult{nil, err}
	}

	model.put = func(id string, item interface{}) modelResult {
		var updateTask Task
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return modelResult{nil, err}
		}
		updateTask = item.(Task)
		updateTask.Id = int64_id
		err = db.Update(updateTask)
		if err != nil {
			return modelResult{nil, err}
		}
		return modelResult{item, nil}
	}

	model.delete = func(id string) modelResult {
		int64_id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return modelResult{nil, err}
		}
		err = db.Delete(Task{Id: int64_id})
		if err != nil {
			return modelResult{nil, err}
		}
		return modelResult{nil, nil}
	}

	return model
}
