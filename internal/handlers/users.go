package handlers

import (
	"net/http"
)

func ShowAllUsersHandler(w http.ResponseWriter, r *http.Request)  {}
func CreateUserHandler(w http.ResponseWriter, r *http.Request)    {}
func ShowUserHandler(w http.ResponseWriter, r *http.Request)      {}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request)    {}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request)    {}
func ShowUserTasksHandler(w http.ResponseWriter, r *http.Request) {}
func SearchUsersHandler(w http.ResponseWriter, r *http.Request)   {}
