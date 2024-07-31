package postgres

import (
	"database/sql"
	"fmt"
	"pm-service/internal/repository/models"
)

type ProjectModel struct {
	DB *sql.DB
}

func (m *ProjectModel) Insert(input *models.ProjectInput) (int, error) {
	var id int
	stmt := `INSERT INTO projects (title, description, manager_id, completed) VALUES ($1, $2, $3, $4) RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Title, input.Description, input.ManagerID, input.Completed).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *ProjectModel) Get(id string) (*models.Project, error) {
	s := &models.Project{}

	stmt := `SELECT id, title, description, manager_id, created, completed FROM projects WHERE id = $1;`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Description, &s.ManagerID, &s.Created, &s.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return s, models.ErrNoRecord
		}

		return s, err
	}

	return s, nil
}

func (m *ProjectModel) Delete(id string) error {
	var row int
	stmt := `DELETE FROM projects WHERE id = $1 RETURNING id;`

	err := m.DB.QueryRow(stmt, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *ProjectModel) Update(id string, input *models.ProjectInput) error {
	var row int
	stmt := `UPDATE projects SET title = $1, description = $2, manager_id = $3, completed = $4 WHERE id = $5 RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Title, input.Description, input.ManagerID, input.Completed, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *ProjectModel) GetAll() ([]*models.Project, error) {
	stmt := `SELECT * FROM projects;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Projects := []*models.Project{}

	for rows.Next() {
		s := &models.Project{}
		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.ManagerID, &s.Created, &s.Completed)
		if err != nil {
			return nil, err
		}
		Projects = append(Projects, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Projects, nil
}

func (m *ProjectModel) GetAllBy(arg, val string) ([]*models.Project, error) {
	stmt := fmt.Sprintf(`SELECT id, title, description, manager_id, created, completed FROM projects WHERE %s = $1;`, arg)

	rows, err := m.DB.Query(stmt, val)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	projects := []*models.Project{}

	for rows.Next() {
		s := &models.Project{}
		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.ManagerID, &s.Created, &s.Completed)
		if err != nil {
			return nil, err
		}
		projects = append(projects, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
