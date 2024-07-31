package mock

import (
	"pm-service/internal/repository/models"
	"time"
)

type UserModel struct {
	DB []*models.User
}

func (m *UserModel) Insert(input *models.UserInput) (int, error) {
	id := len(m.DB) + 1
	m.DB = append(m.DB, &models.User{id, input.Name, input.Email, input.Role, time.Now().Format("2000-01-01")})

	return id, nil
}

func (m *UserModel) Get(id string) (*models.User, error) {
	s := &models.User{}

	return s, nil
}

func (m *UserModel) Delete(id string) error {
	return nil
}

func (m *UserModel) Update(id string, input *models.UserInput) error {
	return nil
}

func (m *UserModel) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (m *UserModel) GetAllBy(arg, val string) ([]*models.User, error) {
	return nil, nil
}
