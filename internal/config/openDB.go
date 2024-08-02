package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/tanimutomo/sqlfile"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func OpenDB(scriptPath string) (*sql.DB, error) {
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", "localhost", "5432", "admin", "password", "database")
	psqlInfo := "postgresql://admin:WIgXi6DPeWA5sCjWxEHFCYwlabZsuNOS@dpg-cqm8b3jqf0us73a76c1g-a/database_54sg"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	/*stmts := []string{
		"CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL, email VARCHAR(100) NOT NULL, role VARCHAR(50) NOT NULL, created VARCHAR(50) DEFAULT CURRENT_DATE);",
		"GRANT ALL PRIVILEGES ON users TO admin;",
		"GRANT ALL PRIVILEGES ON SEQUENCE users_id_seq TO admin;",
		"CREATE TABLE IF NOT EXISTS projects (id SERIAL PRIMARY KEY, title VARCHAR(50) NOT NULL, description VARCHAR(100) NOT NULL, manager_id INTEGER NOT NULL, created VARCHAR(50) DEFAULT CURRENT_DATE, completed VARCHAR(50) NOT NULL, FOREIGN KEY (manager_id) REFERENCES users(id));",
		"GRANT ALL PRIVILEGES ON projects TO admin;",
		"GRANT ALL PRIVILEGES ON SEQUENCE projects_id_seq TO admin;",
		"CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, title VARCHAR(50) NOT NULL, description VARCHAR(100) NOT NULL, priority VARCHAR(50) NOT NULL, status VARCHAR(50) NOT NULL, assignee_id INTEGER NOT NULL, project_id INTEGER NOT NULL, created VARCHAR(50) DEFAULT CURRENT_DATE, completed VARCHAR(50) NOT NULL, FOREIGN KEY (assignee_id) REFERENCES users(id), FOREIGN KEY (project_id) REFERENCES projects(id));",
		"GRANT ALL PRIVILEGES ON tasks TO admin;",
		"GRANT ALL PRIVILEGES ON SEQUENCE tasks_id_seq TO admin;",
	}*/

	// Initialize SqlFile
	s := sqlfile.New()

	// Load input file and store queries written in the file
	err = s.File(scriptPath)
	if err != nil {
		return nil, err
	}

	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	_, err = s.Exec(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
