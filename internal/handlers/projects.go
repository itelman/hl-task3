package handlers

import (
	"encoding/json"
	"net/http"
	"pm-service/internal/handlers/errors"
	"pm-service/internal/service/helpers"
	"strings"

	"github.com/gorilla/mux"
)

// @Summary		List all projects
// @Description	Get a list of all projects
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.Project
// @Failure		500	{object}	map[string]string
// @Router			/projects [get]
func (h *Handler) ShowAllProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := h.projects.GetAll()
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// @Summary		Create a new project
// @Description	Create a new project with the given details
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			project	body		models.ProjectInput	true	"Project details"
// @Success		201		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/projects [post]
func (h *Handler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	input := h.inputs.NewProjectInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	id, err := h.projects.Insert(&input)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	if err := helpers.WriteJSON(w, http.StatusCreated, map[string]interface{}{"id": id}, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}
}

// @Summary		Get project by ID
// @Description	Get a project by its ID
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Project ID"
// @Success		200	{object}	models.Project
// @Failure		500	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/projects/{id} [get]
func (h *Handler) ShowProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	project, err := h.projects.Get(id)
	if err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// @Summary		Update project details
// @Description	Update project details by ID
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"Project ID"
// @Param			project	body		models.ProjectInput	true	"Project details"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		404		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/projects/{id} [put]
func (h *Handler) UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	input := h.inputs.NewProjectInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	if err := h.projects.Update(id, &input); err != nil {
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

// @Summary		Delete project by ID
// @Description	Delete a project by its ID
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Project ID"
// @Success		200	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/projects/{id} [delete]
func (h *Handler) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.projects.Delete(id); err != nil {
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

// @Summary		Get tasks by project ID
// @Description	Get tasks associated with a project by project ID
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Project ID"
// @Success		200	{array}		models.Task
// @Failure		404	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/projects/{id}/tasks [get]
func (h *Handler) ShowProjectTasksHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := h.projects.Get(id)
	if err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	tasks, err := h.tasks.GetAllBy("project_id", id)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// @Summary		Search projects by query
// @Description	Search projects by title or manager ID
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			title	query		string	false	"Project title"
// @Param			manager	query		string	false	"Manager ID"
// @Success		200		{array}		models.Project
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/projects/search [get]
func (h *Handler) SearchProjectsHandler(w http.ResponseWriter, r *http.Request) {
	var query string

	if r.URL.Query().Get("title") != "" {
		query = "title"
	} else if r.URL.Query().Get("manager") != "" {
		query = "manager_id"
	}

	queryURL := strings.ReplaceAll(query, "_id", "")

	if query == "" || r.URL.Query().Get(queryURL) == "" {
		errors.BadRequestResponse(w, r)
		return
	}

	projects, err := h.projects.GetAllBy(query, r.URL.Query().Get(queryURL))
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
