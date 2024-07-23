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

	router.HandleFunc("/users", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/users", handlers.DeleteTaskHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.DeleteTaskHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/users/{id}/tasks", handlers.ShowTaskHandler).Methods("GET")
	router.HandleFunc("/users/search", handlers.DeleteTaskHandler).Methods("GET")

	router.HandleFunc("/tasks", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", handlers.DeleteTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.ShowTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/tasks/search", handlers.DeleteTaskHandler).Methods("GET")

	router.HandleFunc("/projects", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/projects", handlers.DeleteTaskHandler).Methods("POST")
	router.HandleFunc("/projects/{id}", handlers.DeleteTaskHandler).Methods("GET")
	router.HandleFunc("/projects/{id}", handlers.DeleteTaskHandler).Methods("PUT")
	router.HandleFunc("/projects/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/projects/{id}/tasks", handlers.ShowTaskHandler).Methods("GET")
	router.HandleFunc("/projects/search", handlers.DeleteTaskHandler).Methods("GET")

	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(handlers.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(handlers.MethodNotAllowedResponse)

	return standardMiddleware.Then(router)
}
