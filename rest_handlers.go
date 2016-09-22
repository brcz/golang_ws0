package main

//Step4: Implement CRUD handlers

import (
	"net/http"

	"github.com/abiosoft/river"
)
// swagger:route GET /todo/:id getTODORecord Task getTODORecord 
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


// getTODORecord handles GET /todo/:id.
func getTODORecord(c *river.Context, model Model) {
	task := model.get(c.Param("id"))
	if task == nil {
		c.RenderEmpty(http.StatusNotFound)
		return
	}
	c.Render(http.StatusOK, task)
}

// getTODOList handles GET /todo.
func getTODOList(c *river.Context, model Model) {
	c.Render(http.StatusOK, model.getAll())
}

// addTODORecord handles POST /todo.
func addTODORecord(c *river.Context, model Model) {
	var tasks TaskList
	if err := c.DecodeJSONBody(&tasks); err != nil {
		c.Render(http.StatusBadRequest, err)
		return
	}
	for i := range tasks {
		model.add(tasks[i])
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

	model.put(id, task)
	c.Render(http.StatusOK, task)
}

// deleteTODORecord handles DELETE /todo/:id.
func deleteTODORecord(c *river.Context, model Model) {
	model.delete(c.Param("id"))
	c.RenderEmpty(http.StatusNoContent)
}
