package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID      int
	Name    string
	Email   string
	Role    string
	Created time.Time
}

type Task struct {
	ID           int
	Name         string
	Description  string
	Priority     string
	Status       string
	SupervisorID int
	ProjectID    int
	Created      time.Time
	Completed    time.Time
}

type Project struct {
	ID          int
	Name        string
	Description string
	ManagerID   int
	Created     time.Time
	Completed   time.Time
}
