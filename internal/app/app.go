package app

import (
	"database/sql"
	"log"
	"net/http"
	"pm-service/internal/config"
	"pm-service/internal/handlers"
)

type Application struct {
	Port   int
	DB     *sql.DB
	Routes http.Handler
}

func NewApp(port int, scriptPath string) *Application {
	db, err := config.OpenDB(scriptPath)
	if err != nil {
		log.Fatalln(err)
	}

	handlers := handlers.New(db)

	return &Application{port, db, config.Routing(handlers)}
}
