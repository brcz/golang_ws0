package main

//Step4: Implement CRUD handlers

import (
	"encoding/json"
	"fmt"
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

// getTODORecord handles GET /todo/:id.
func getTODORecord(c *river.Context, model Model) {
	modelData := model.get(c.Param("id"))
	if modelData.status != nil {
		c.Render(http.StatusNotFound, modelData.status)
	} else {
		c.Render(http.StatusOK, modelData.result)
	}
}

func getTODORecordExt(model Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modelData := model.get(r.FormValue("id"))
		if modelData.status != nil {
			//c.Render(http.StatusNotFound, modelData.status)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", modelData.status)))
		} else {
			//c.Render(http.StatusOK, modelData.result)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(modelData.result)
		}
	}
}

// getTODOList handles GET /todo.
func getTODOList(c *river.Context, model Model) {
	modelData := model.getAll()
	if modelData.status != nil {
		c.Render(http.StatusNoContent, modelData.status)
	} else {
		c.Render(http.StatusOK, modelData.result)
	}
}

// getTODOList wrapped to use in classic mode
func getTODOListExt(model Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modelData := model.getAll()
		if modelData.status != nil {
			//c.Render(http.StatusNoContent, modelData.status)
			w.WriteHeader(http.StatusNoContent)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", modelData.status)))
		} else {
			//c.Render(http.StatusOK, modelData.result)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(modelData.result)
		}
	}
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
func addTODORecordExt(model Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var tasks TaskList

		if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		for i := range tasks {
			modelData := model.add(tasks[i])
			if modelData.status != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set("Content-Type", "text/plain")
				w.Write([]byte(fmt.Sprintf("%s", modelData.status)))
				//c.Render(http.StatusInternalServerError, modelData.status)
				return
			}
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

		//c.Render(http.StatusCreated, tasks)
	}
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

func updateTODORecordExt(model Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		var task Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		modelData := model.put(id, task)
		if modelData.status != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", modelData.status)))
			//c.Render(http.StatusInternalServerError, modelData.status)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(modelData.result)

		//c.Render(http.StatusOK, task)

	}
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
func deleteTODORecordExt(model Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modelData := model.delete(r.FormValue("id"))
		if modelData.status != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Sprintf("%s", modelData.status)))
			//c.Render(http.StatusInternalServerError, modelData.status)
		} else {
			w.WriteHeader(http.StatusNoContent)
			//c.RenderEmpty(http.StatusNoContent)
		}
	}
}
