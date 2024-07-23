package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var store sync.Map

// deleteTaskHandler godoc
//
//	@Summary		Delete a task
//	@Description	Delete a task by ID
//	@Tags			tasks
//	@Param			id	path	string	true	"Task ID"
//	@Success		204
//	@Failure		404	{string}	string	"Not Found"
//	@Router			/api/todo-list/tasks/{id} [delete]
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if _, ok := store.Load(id); !ok {
		NotFoundResponse(w, r)
		return
	}

	store.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

func ShowTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	val, ok := store.Load(id)
	if !ok {
		NotFoundResponse(w, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(val)
}
