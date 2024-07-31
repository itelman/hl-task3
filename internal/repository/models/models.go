package models

type User struct {
	ID      int
	Name    string
	Email   string
	Role    string
	Created string
}

type Task struct {
	ID          int
	Title       string
	Description string
	Priority    string
	Status      string
	AssigneeID  int
	ProjectID   int
	Created     string
	Completed   string
}

type Project struct {
	ID          int
	Title       string
	Description string
	ManagerID   int
	Created     string
	Completed   string
}
