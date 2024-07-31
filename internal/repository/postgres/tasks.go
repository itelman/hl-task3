package postgres

import (
	"database/sql"
	"fmt"
	"pm-service/internal/repository/models"
)

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Insert(input *models.TaskInput) (int, error) {
	var id int
	stmt := `INSERT INTO tasks (title, description, priority, status, assignee_id, project_id, completed) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Title, input.Description, input.Priority, input.Status, input.AssigneeID, input.ProjectID, input.Completed).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *TaskModel) Get(id string) (*models.Task, error) {
	s := &models.Task{}

	stmt := `SELECT id, title, description, priority, status, assignee_id, project_id, created, completed FROM tasks WHERE id = $1;`
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Description, &s.Priority, &s.Status, &s.AssigneeID, &s.ProjectID, &s.Created, &s.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return s, models.ErrNoRecord
		}

		return s, err
	}

	return s, nil
}

func (m *TaskModel) Delete(id string) error {
	var row int
	stmt := `DELETE FROM tasks WHERE id = $1 RETURNING id;`

	err := m.DB.QueryRow(stmt, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *TaskModel) Update(id string, input *models.TaskInput) error {
	var row int
	stmt := `UPDATE tasks SET title = $1, description = $2, priority = $3, status = $4, assignee_id = $5, project_id = $6, completed = $7 WHERE id = $8 RETURNING id;`

	err := m.DB.QueryRow(stmt, input.Title, input.Description, input.Priority, input.Status, input.AssigneeID, input.ProjectID, input.Completed, id).Scan(&row)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.ErrNoRecord
		}

		return err
	}

	return nil
}

func (m *TaskModel) GetAll() ([]*models.Task, error) {
	stmt := `SELECT * FROM tasks;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []*models.Task{}

	for rows.Next() {
		s := &models.Task{}
		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.Priority, &s.Status, &s.AssigneeID, &s.ProjectID, &s.Created, &s.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskModel) GetAllBy(arg, val string) ([]*models.Task, error) {
	stmt := fmt.Sprintf(`SELECT id, title, description, priority, status, assignee_id, project_id, created, completed FROM tasks WHERE %s = $1;`, arg)

	rows, err := m.DB.Query(stmt, val)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := []*models.Task{}

	for rows.Next() {
		s := &models.Task{}
		err = rows.Scan(&s.ID, &s.Title, &s.Description, &s.Priority, &s.Status, &s.AssigneeID, &s.ProjectID, &s.Created, &s.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
