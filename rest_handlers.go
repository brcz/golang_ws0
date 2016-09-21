package main

//Step4: Implement CRUD handlers

import (
	"net/http"

	"github.com/abiosoft/river"
)

// getUser handles GET /user/:id.
func getTODORecord(c *river.Context, model Model) {
	task := model.get(c.Param("id"))
	if task == nil {
		c.RenderEmpty(http.StatusNotFound)
		return
	}
	c.Render(http.StatusOK, task)
}

// getAllUser handles GET /user.
func getTODOList(c *river.Context, model Model) {
	c.Render(http.StatusOK, model.getAll())
}

// addUser handles POST /user.
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

// updateUser handles PUT /user/:id.
func updateTODORecord(c *river.Context, model Model) {
	id := c.Param("id")
	var task Task
	if err := c.DecodeJSONBody(&task); err != nil {
		c.Render(http.StatusBadRequest, err)
		return
	}
	_ = id
	//model.put(id, user)
	c.Render(http.StatusOK, task)
}

// deleteUser handles DELETE /user/:id.
func deleteTODORecord(c *river.Context, model Model) {
	//model.delete(c.Param("id"))
	c.RenderEmpty(http.StatusNoContent)
}
