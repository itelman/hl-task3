package mock

import (
	"pm-service/internal/repository/models"
)

type ProjectModel struct {
	DB []*models.Project
}

func (m *ProjectModel) Insert(input *models.ProjectInput) (int, error) {
	var id int

	return id, nil
}

func (m *ProjectModel) Get(id string) (*models.Project, error) {
	s := &models.Project{}

	return s, nil
}

func (m *ProjectModel) Delete(id string) error {

	return nil
}

func (m *ProjectModel) Update(id string, input *models.ProjectInput) error {

	return nil
}

func (m *ProjectModel) GetAll() ([]*models.Project, error) {

	projects := []*models.Project{}

	return projects, nil
}

func (m *ProjectModel) GetAllBy(arg, val string) ([]*models.Project, error) {

	projects := []*models.Project{}

	return projects, nil
}
