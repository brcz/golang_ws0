package main

//Step4: Implement CRUD handlers

import (
	"net/http"

	"github.com/abiosoft/river"
)
// swagger:route GET /:id getTODORecord Task getTODORecord
//
// Fetch TODO task by id or alias.
//
// This will show task by given id or alias.
//
//     Consumes:
//     - application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       default: notFound
//       200: task
type renderInerface interface {
    Render(int, interface{})
    DecodeJSONBody(*interface{}) error
}


// getTODORecord handles GET /todo/:id.
func getTODORecord(c *river.Context, model Model) {
	modelData := model.get(c.Param("id"))
	if modelData.status != nil {
		c.Render(http.StatusNotFound, modelData.status)
	} else {
		c.Render(http.StatusOK, modelData.result)
	}
}

// getTODOList handles GET /todo.
//func getTODOList(c renderInerface, model Model) {
func getTODOList(c *river.Context, model Model) {
	modelData := model.getAll()
	if modelData.status != nil {
		c.Render(http.StatusNoContent, modelData.status)
	} else {
		c.Render(http.StatusOK, modelData.result)
	}
}

func getTODOListExt(c *river.Context, model Model) {
	//c.Render(http.StatusOK, model.getAll())
}

// addTODORecord handles POST /todo.
func addTODORecord(c *river.Context, model Model) {
	var tasks TaskList
	if err := c.DecodeJSONBody(&tasks); err != nil {
		c.Render(http.StatusBadRequest, err)
		return
	}
	for i := range tasks {
		modelData := model.add(tasks[i])
		if modelData.status != nil {
			c.Render(http.StatusInternalServerError, modelData.status)
			return
		}
	}
	c.Render(http.StatusCreated, tasks)
}

// updateTODORecord handles PUT /todo/:id.
func updateTODORecord(c *river.Context, model Model) {
	id := c.Param("id")
	var task Task
	if err := c.DecodeJSONBody(&task); err != nil {
		c.Render(http.StatusBadRequest, err)
		return
	}

	modelData := model.put(id, task)
	if modelData.status != nil {
		c.Render(http.StatusInternalServerError, modelData.status)
		return
	}
	c.Render(http.StatusOK, task)
}

// deleteTODORecord handles DELETE /todo/:id.
func deleteTODORecord(c *river.Context, model Model) {
	modelData := model.delete(c.Param("id"))
	if modelData.status != nil {
		c.Render(http.StatusInternalServerError, modelData.status)
	} else {
		c.RenderEmpty(http.StatusNoContent)
	}
}
