package mock

import (
	"pm-service/internal/repository/models"
)

type TaskModel struct {
	DB []*models.Task
}

func (m *TaskModel) Insert(input *models.TaskInput) (int, error) {
	var id int

	return id, nil
}

func (m *TaskModel) Get(id string) (*models.Task, error) {
	s := &models.Task{}

	return s, nil
}

func (m *TaskModel) Delete(id string) error {
	return nil
}

func (m *TaskModel) Update(id string, input *models.TaskInput) error {
	return nil
}

func (m *TaskModel) GetAll() ([]*models.Task, error) {
	return nil, nil
}

func (m *TaskModel) GetAllBy(arg, val string) ([]*models.Task, error) {
	return nil, nil
}
