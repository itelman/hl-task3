package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pm-service/internal/config"
	"pm-service/internal/handlers"
	"pm-service/internal/repository/models"
	"testing"
)

var router = config.Routing(handlers.Mock())

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		method   string
		input    models.UserInput
		wantCode int
	}{
		{
			name:     "test1",
			path:     "/invalid-path",
			method:   "POST",
			input:    models.UserInput{Name: "alice", Email: "alice02@mail.com", Role: "manager"},
			wantCode: http.StatusNotFound,
		},
		{
			name:     "test2",
			path:     "/users/invalid-path",
			method:   "POST",
			input:    models.UserInput{Name: "alice", Email: "alice02@mail.com", Role: "manager"},
			wantCode: http.StatusNotFound,
		},
		{
			name:     "test3",
			path:     "/users",
			method:   "PUT",
			input:    models.UserInput{Name: "alice", Email: "alice02@mail.com", Role: "manager"},
			wantCode: http.StatusMethodNotAllowed,
		},
		{
			name:     "test4",
			path:     "/users",
			method:   "POST",
			input:    models.UserInput{},
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "test5",
			path:     "/users",
			method:   "POST",
			input:    models.UserInput{Name: "alice", Email: "alice02@mail.", Role: "manager"},
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "test6",
			path:     "/users",
			method:   "POST",
			input:    models.UserInput{Name: "alice", Email: "alice02@mail.com", Role: "manager"},
			wantCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatal(err)
			}

			req, err := http.NewRequest(tt.method, tt.path, bytes.NewReader(json))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if rr.Code != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, tt.wantCode)
			}
		})
	}
}
