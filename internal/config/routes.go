package config

import (
	"net/http"
	"pm-service/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "pm-service/docs"
)

//	@title			Todo List API
//	@version		1.0
//	@description	This is a simple Todo List API.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @BasePath
func Routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our app receives.
	standardMiddleware := alice.New(recoverPanic, logRequest, secureHeaders)

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.ShowAllProjectsHandler).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.ShowUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/users/{id}/tasks", handlers.ShowUserTasksHandler).Methods("GET")
	router.HandleFunc("/users/search", handlers.SearchUsersHandler).Methods("GET")

	router.HandleFunc("/tasks", handlers.ShowAllTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.ShowTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tasks/search", handlers.SearchTasksHandler).Methods("GET")

	router.HandleFunc("/projects", handlers.ShowAllProjectsHandler).Methods("GET")
	router.HandleFunc("/projects", handlers.CreateProjectHandler).Methods("POST")
	router.HandleFunc("/projects/{id}", handlers.ShowProjectHandler).Methods("GET")
	router.HandleFunc("/projects/{id}", handlers.UpdateProjectHandler).Methods("PUT")
	router.HandleFunc("/projects/{id}", handlers.DeleteProjectHandler).Methods("DELETE")
	router.HandleFunc("/projects/{id}/tasks", handlers.ShowProjectTasksHandler).Methods("GET")
	router.HandleFunc("/projects/search", handlers.SearchProjectsHandler).Methods("GET")

	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedResponse)

	return standardMiddleware.Then(router)
}
