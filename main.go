package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"pm-service/internal/app"
)

func main() {
	port := flag.Int("port", 8080, "port for api")
	flag.Parse()

	app := app.NewApp(*port)

	defer app.DB.Close()

	fmt.Printf("Server starting on http://localhost:%d\n\n", app.Port)
	err := http.ListenAndServe(":8080", app.Routes)
	if err != nil {
		log.Fatalln(err)
	}
}
