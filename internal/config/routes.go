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

//	@BasePath	/health
func Routes(handlers *handlers.Handler) http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our app receives.
	standardMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.ShowAllUsersHandler).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/search", handlers.SearchUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.ShowUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users/{id}/tasks", handlers.ShowUserTasksHandler).Methods("GET")

	router.HandleFunc("/tasks", handlers.ShowAllTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/search", handlers.SearchTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.ShowTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")

	router.HandleFunc("/projects", handlers.ShowAllProjectsHandler).Methods("GET")
	router.HandleFunc("/projects", handlers.CreateProjectHandler).Methods("POST")
	router.HandleFunc("/projects/search", handlers.SearchProjectsHandler).Methods("GET")
	router.HandleFunc("/projects/{id}", handlers.ShowProjectHandler).Methods("GET")
	router.HandleFunc("/projects/{id}", handlers.UpdateProjectHandler).Methods("PUT")
	router.HandleFunc("/projects/{id}", handlers.DeleteProjectHandler).Methods("DELETE")
	router.HandleFunc("/projects/{id}/tasks", handlers.ShowProjectTasksHandler).Methods("GET")

	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(errors.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(errors.MethodNotAllowedResponse)

	return standardMiddleware.Then(router)
}
