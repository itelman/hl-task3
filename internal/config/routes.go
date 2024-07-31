package config

import (
	"net/http"
	"pm-service/internal/handlers"
	"pm-service/internal/handlers/errors"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "pm-service/docs"
)

type routes struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

//	@title			Project Management Service
//	@version		1.0
//	@description	This is an API server for project management service.
//	@host			localhost:8080
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @BasePath	/health
func Routing(handlers *handlers.Handler) http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our app receives.
	standardMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)

	router := mux.NewRouter()

	routes := []routes{
		{"/health", handlers.HealthCheckHandler, http.MethodGet},
		{"/users", handlers.ShowAllUsersHandler, http.MethodGet},
		{"/users", handlers.CreateUserHandler, http.MethodPost},
		{"/users/search", handlers.SearchUsersHandler, http.MethodGet},
		{"/tasks", handlers.ShowAllTasksHandler, http.MethodGet},
		{"/tasks", handlers.CreateTaskHandler, http.MethodPost},
		{"/tasks/search", handlers.SearchTasksHandler, http.MethodGet},
		{"/projects", handlers.ShowAllProjectsHandler, http.MethodGet},
		{"/projects", handlers.CreateProjectHandler, http.MethodPost},
		{"/projects/search", handlers.SearchProjectsHandler, http.MethodGet},
		{"/users/{id:[0-9]+}", handlers.ShowUserHandler, http.MethodGet},
		{"/users/{id:[0-9]+}", handlers.UpdateUserHandler, http.MethodPut},
		{"/users/{id:[0-9]+}", handlers.DeleteUserHandler, http.MethodDelete},
		{"/users/{id:[0-9]+}/tasks", handlers.ShowUserTasksHandler, http.MethodGet},
		{"/projects/{id:[0-9]+}", handlers.ShowProjectHandler, http.MethodGet},
		{"/projects/{id:[0-9]+}", handlers.UpdateProjectHandler, http.MethodPut},
		{"/projects/{id:[0-9]+}", handlers.DeleteProjectHandler, http.MethodDelete},
		{"/projects/{id:[0-9]+}/tasks", handlers.ShowProjectTasksHandler, http.MethodGet},
		{"/tasks/{id:[0-9]+}", handlers.ShowTaskHandler, http.MethodGet},
		{"/tasks/{id:[0-9]+}", handlers.UpdateTaskHandler, http.MethodPut},
		{"/tasks/{id:[0-9]+}", handlers.DeleteTaskHandler, http.MethodDelete},
	}

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(errors.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(errors.MethodNotAllowedResponse)

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return standardMiddleware.Then(router)
}
