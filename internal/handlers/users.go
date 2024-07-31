package handlers

import (
	"encoding/json"
	"net/http"
	"pm-service/internal/handlers/errors"
	"pm-service/internal/service/helpers"

	"github.com/gorilla/mux"
)

// @Summary		List all users
// @Description	Get a list of all users
// @Tags			Users
// @Accept			json
// @Produce		json
// @Success		200	{array}		models.User
// @Failure		500	{object}	map[string]string
// @Router			/users [get]
func (h *Handler) ShowAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.users.GetAll()
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// @Summary		Create a new user
// @Description	Create a new user with the given details
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			user	body		models.UserInput	true	"User details"
// @Success		201		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/users [post]
func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	input := h.inputs.NewUserInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	id, err := h.users.Insert(&input)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	if err := helpers.WriteJSON(w, http.StatusCreated, map[string]interface{}{"id": id}, nil); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}
}

// @Summary		Get user by ID
// @Description	Get a user by their ID
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	models.User
// @Failure		500	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Router			/users/{id} [get]
func (h *Handler) ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := h.users.Get(id)
	if err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary		Update user details
// @Description	Update user details by ID
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"User ID"
// @Param			user	body		models.UserInput	true	"User details"
// @Success		200		{object}	map[string]string
// @Failure		400		{object}	map[string]string
// @Failure		404		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/users/{id} [put]
func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	input := h.inputs.NewUserInput()

	if err := helpers.ReadJSON(w, r, &input); err != nil {
		errors.BadRequestResponse(w, r)
		return
	}

	if !input.IsValid() {
		errors.BadRequestResponse(w, r)
		return
	}

	if err := h.users.Update(id, &input); err != nil {
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

// @Summary		Delete user by ID
// @Description	Delete a user by their ID
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	map[string]string
// @Failure		404	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/users/{id} [delete]
func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.users.Delete(id); err != nil {
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

// @Summary		Get user tasks
// @Description	Get all tasks assigned to a user
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{array}		models.Task
// @Failure		404	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/users/{id}/tasks [get]
func (h *Handler) ShowUserTasksHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := h.users.Get(id)
	if err != nil {
		if err == h.errors.NoRecordError() {
			errors.NotFoundResponse(w, r)
		} else {
			errors.ServerErrorResponse(w, r, err)
		}
		return
	}

	tasks, err := h.tasks.GetAllBy("assignee_id", id)
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// @Summary		Search users by name or email
// @Description	Search users by name or email
// @Tags			Users
// @Accept			json
// @Produce		json
// @Param			name	query		string	false	"User name"
// @Param			email	query		string	false	"User email"
// @Success		200		{array}		models.User
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/users/search [get]
func (h *Handler) SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	var query string

	if r.URL.Query().Get("name") != "" {
		query = "name"
	} else if r.URL.Query().Get("email") != "" {
		query = "email"
	}

	if query == "" || r.URL.Query().Get(query) == "" {
		errors.BadRequestResponse(w, r)
		return
	}

	users, err := h.users.GetAllBy(query, r.URL.Query().Get(query))
	if err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}
