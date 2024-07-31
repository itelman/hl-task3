package handlers

import (
	"encoding/json"
	"net/http"
	"pm-service/internal/handlers/errors"
	"pm-service/internal/service/helpers"
	"strings"

	"github.com/gorilla/mux"
)

// @Summary		List all tasks
// @Description	Get a list of all tasks
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.Task
// @Failure		500	{object}	map[string]string
// @Router			/tasks [get]
func (h *Handler) ShowAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.tasks.GetAll()
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// @Summary		Create a new task
// @Description	Create a new task with the given details
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			task	body		models.TaskInput	true	"Task details"
// @Success		201		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/tasks [post]
func (h *Handler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	input := h.inputs.NewTaskInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	id, err := h.tasks.Insert(&input)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	if err := helpers.WriteJSON(w, http.StatusCreated, map[string]interface{}{"id": id}, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}
}

// @Summary		Get task by ID
// @Description	Get a task by its ID
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Task ID"
// @Success		200	{object}	models.Task
// @Failure		500	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/tasks/{id} [get]
func (h *Handler) ShowTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, err := h.tasks.Get(id)
	if err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// @Summary		Update task details
// @Description	Update task details by ID
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Task ID"
// @Param			task	body		models.TaskInput	true	"Task details"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		404		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/tasks/{id} [put]
func (h *Handler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	input := h.inputs.NewTaskInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	if err := h.tasks.Update(id, &input); err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := helpers.WriteJSON(w, http.StatusOK, map[string]interface{}{"status": "OK"}, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}
}

// @Summary		Delete task by ID
// @Description	Delete a task by its ID
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Task ID"
// @Success		200	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/tasks/{id} [delete]
func (h *Handler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.tasks.Delete(id); err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := helpers.WriteJSON(w, http.StatusOK, map[string]interface{}{"status": "OK"}, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}
}

// @Summary		Search tasks by query
// @Description	Search tasks by title, status, priority, assignee, or project
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			title		query		string	false	"Task title"
// @Param			status		query		string	false	"Task status"
// @Param			priority	query		string	false	"Task priority"
// @Param			assignee	query		string	false	"Task assignee"
// @Param			project		query		string	false	"Project ID"
// @Success		200			{array}		models.Task
// @Failure		400			{object}	map[string]string
// @Failure		500			{object}	map[string]string
// @Router			/tasks/search [get]
func (h *Handler) SearchTasksHandler(w http.ResponseWriter, r *http.Request) {
	var query string

	if r.URL.Query().Get("title") != "" {
		query = "title"
	} else if r.URL.Query().Get("status") != "" {
		query = "status"
	} else if r.URL.Query().Get("priority") != "" {
		query = "priority"
	} else if r.URL.Query().Get("assignee") != "" {
		query = "assignee_id"
	} else if r.URL.Query().Get("project") != "" {
		query = "project_id"
	}

	queryURL := strings.ReplaceAll(query, "_id", "")

	if query == "" || r.URL.Query().Get(queryURL) == "" {
		errors.BadRequestResponse(w, r)
		return
	}

	tasks, err := h.tasks.GetAllBy(query, r.URL.Query().Get(queryURL))
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
