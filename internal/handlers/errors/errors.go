package errors

import (
	"fmt"
	"log"
	"net/http"
	"pm-service/internal/service/helpers"
	"runtime/debug"
)

func errorResponse(w http.ResponseWriter, status int, message interface{}) {
	env := map[string]interface{}{"error": message}

	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	if err := log.Output(2, trace); err != nil {
		log.Println("issue with printing error logs", err)
	}

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid request"
	errorResponse(w, http.StatusBadRequest, message)
}
