package postgres

import (
	"database/sql"
	"fmt"
	"pm-service/internal/repository/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(input *models.UserInput) (int, error) {
	var id int
	stmt := `INSERT INTO users (name, email, role) VALUES ($1, $2, $3) RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Name, input.Email, input.Role).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *UserModel) Get(id string) (*models.User, error) {
	s := &models.User{}

	stmt := `SELECT id, name, email, role, created FROM users WHERE id = $1;`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Name, &s.Email, &s.Role, &s.Created)
	if err != nil {
		if err == sql.ErrNoRows {
			return s, models.ErrNoRecord
		}

		return s, err
	}

	return s, nil
}

func (m *UserModel) Delete(id string) error {
	var row int
	stmt := `DELETE FROM users WHERE id = $1 RETURNING id;`

	err := m.DB.QueryRow(stmt, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *UserModel) Update(id string, input *models.UserInput) error {
	var row int
	stmt := `UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4 RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Name, input.Email, input.Role, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *UserModel) GetAll() ([]*models.User, error) {
	stmt := `SELECT * FROM users;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*models.User{}

	for rows.Next() {
		s := &models.User{}
		err = rows.Scan(&s.ID, &s.Name, &s.Email, &s.Role, &s.Created)
		if err != nil {
			return nil, err
		}
		users = append(users, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *UserModel) GetAllBy(arg, val string) ([]*models.User, error) {
	stmt := fmt.Sprintf(`SELECT id, name, email, role, created FROM users WHERE %s = $1;`, arg)

	rows, err := m.DB.Query(stmt, val)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []*models.User{}

	for rows.Next() {
		s := &models.User{}
		err = rows.Scan(&s.ID, &s.Name, &s.Email, &s.Role, &s.Created)
		if err != nil {
			return nil, err
		}
		users = append(users, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
