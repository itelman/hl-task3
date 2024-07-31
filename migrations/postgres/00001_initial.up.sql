CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created VARCHAR(50) DEFAULT CURRENT_DATE
);

GRANT ALL PRIVILEGES ON users TO admin;
GRANT ALL PRIVILEGES ON SEQUENCE users_id_seq TO admin;

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    manager_id INTEGER NOT NULL,
    created VARCHAR(50) DEFAULT CURRENT_DATE,
    completed VARCHAR(50) NOT NULL,
    FOREIGN KEY (manager_id) REFERENCES users(id)
);

GRANT ALL PRIVILEGES ON projects TO admin;
GRANT ALL PRIVILEGES ON SEQUENCE projects_id_seq TO admin;

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    priority VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    assignee_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    created VARCHAR(50) DEFAULT CURRENT_DATE,
    completed VARCHAR(50) NOT NULL,
    FOREIGN KEY (assignee_id) REFERENCES users(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);

GRANT ALL PRIVILEGES ON tasks TO admin;
GRANT ALL PRIVILEGES ON SEQUENCE tasks_id_seq TO admin;