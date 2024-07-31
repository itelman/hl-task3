package models

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	dateRegex  = regexp.MustCompile("^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$")
	priorRegex = regexp.MustCompile("^(high|medium|low)$")
	idRegex    = regexp.MustCompile("^([0-9]+)$")
	EmailRX    = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	statusRX   = regexp.MustCompile("^(to do|in progress|completed)$")
)

type Input struct {
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
	AssigneeID  int    `json:"assignee_id"`
	ProjectID   int    `json:"project_id"`
	Completed   string `json:"completed"`
}

type ProjectInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ManagerID   int    `json:"manager_id"`
	Completed   string `json:"completed"`
}

func (i *Input) NewUserInput() UserInput {
	return UserInput{}
}

func (i *Input) NewTaskInput() TaskInput {
	return TaskInput{}
}

func (i *Input) NewProjectInput() ProjectInput {
	return ProjectInput{}
}

func (i *ProjectInput) IsValid() bool {
	return dateRegex.MatchString(i.Completed) || idRegex.MatchString(strconv.Itoa(i.ManagerID))
}

func (i *TaskInput) IsValid() bool {
	return dateRegex.MatchString(i.Completed) || priorRegex.MatchString(strings.ToLower(i.Priority)) || statusRX.MatchString(strings.ToLower(i.Status)) || idRegex.MatchString(strconv.Itoa(i.AssigneeID)) || idRegex.MatchString(strconv.Itoa(i.ProjectID))
}

func (i *UserInput) IsValid() bool {
	return EmailRX.MatchString(i.Email)
}
