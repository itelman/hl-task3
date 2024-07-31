# Project Management Service

## Introduction

This project is a RESTful API for a Project Management microservice built in Go. It provides endpoints for managing users, tasks, and projects, including operations such as creating, updating, deleting, and searching.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/itelman/hl-task3.git
   cd hl-task3
   ```

2. Run the server:

   ```sh
   go run main.go
   ```

Or use the link: https://hl-task3.onrender.com

## API Endpoints

### Users
#### URL: /users

- **GET /users**: Get a list of all users.

- **POST /users**: Create a new user.
  - Request Body:
    ```json
    {
        "name": "John Doe",
        "email": "johndoe@example.com",
        "role": "admin"
    }
    ```

- **GET /users/{id}**: Get details of a specific user.

- **PUT /users/{id}**: Update details of a specific user.
  - Request Body:
    ```json
    {
        "name": "John Doe",
        "email": "johndoe@example.com",
        "role": "manager"
    }
    ```

- **DELETE /users/{id}**: Delete a specific user.

### Projects
#### URL: /projects

- **GET /projects**: Get a list of all projects.

- **POST /projects**: Create a new project.
  - Request Body:
    ```json
    {
        "title": "Project Alpha",
        "description": "A new innovative project",
        "manager_id": 1,
        "completed": "2024-12-31"
    }
    ```

- **GET /projects/{id}**: Get details of a specific project.

- **PUT /projects/{id}**: Update details of a specific project.
  - Request Body:
    ```json
    {
        "title": "Project Beta",
        "description": "An updated innovative project",
        "manager_id": 1,
        "completed": "2024-12-31"
    }
    ```

- **DELETE /projects/{id}**: Delete a specific project.

### Tasks
#### URL: /tasks

- **GET /tasks**: Get a list of all tasks.

- **POST /tasks**: Create a new task.
  - Request Body:
    ```json
    {
        "title": "Finish Report",
        "description": "Complete the quarterly financial report",
        "priority": "High",
        "status": "Completed",
        "assignee_id": 3,
        "project_id": 5,
        "completed": "2024-07-10"
    }
    ```

- **GET /tasks/{id}**: Get details of a specific task.

- **PUT /tasks/{id}**: Update details of a specific task.
  - Request Body:
    ```json
    {
        "title": "Finish Report",
        "description": "Complete the quarterly financial report and review",
        "priority": "High",
        "status": "Completed",
        "assignee_id": 3,
        "project_id": 5,
        "completed": "2024-07-10"
    }
    ```

- **DELETE /tasks/{id}**: Delete a specific task.