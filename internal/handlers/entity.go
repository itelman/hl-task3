package handlers

import (
	"database/sql"
	"pm-service/internal/repository/mock"
	"pm-service/internal/repository/models"
	"pm-service/internal/repository/postgres"
)

type Handler struct {
	input interface {
		NewUserInput() models.UserInput
		NewTaskInput() models.TaskInput
		NewProjectInput() models.ProjectInput
	}
	errors interface {
		NoRecordError() error
	}
	users interface {
		Insert(*models.UserInput) (int, error)
		Get(string) (*models.User, error)
		Delete(string) error
		Update(string, *models.UserInput) error
		GetAll() ([]*models.User, error)
		GetAllBy(string, string) ([]*models.User, error)
	}
	projects interface {
		Insert(*models.ProjectInput) (int, error)
		Get(string) (*models.Project, error)
		Delete(string) error
		Update(string, *models.ProjectInput) error
		GetAll() ([]*models.Project, error)
		GetAllBy(string, string) ([]*models.Project, error)
	}
	tasks interface {
		Insert(*models.TaskInput) (int, error)
		Get(string) (*models.Task, error)
		Delete(string) error
		Update(string, *models.TaskInput) error
		GetAll() ([]*models.Task, error)
		GetAllBy(string, string) ([]*models.Task, error)
	}
}

func New(db *sql.DB) *Handler {
	return &Handler{
		&models.Input{},
		&models.Errors{},
		&postgres.UserModel{DB: db},
		&postgres.ProjectModel{DB: db},
		&postgres.TaskModel{DB: db},
	}
}

func Mock() *Handler {
	return &Handler{
		&models.Input{},
		&models.Errors{},
		&mock.UserModel{DB: make([]*models.User, 0)},
		&mock.ProjectModel{DB: make([]*models.Project, 0)},
		&mock.TaskModel{DB: make([]*models.Task, 0)},
	}
}
