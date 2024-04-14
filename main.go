package main

import (
	"course-backend/pkg/database"
	http_server "course-backend/pkg/delivery/http"
	"course-backend/pkg/utils"
	"log"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.NewDatabase("main")

	if err != nil {
		return
	}

	go utils.RunWatcher(db)

	http_server := http_server.HttpServer{}
	http_server.SetupDB(db)
	http_server.RouteAll()
	http_server.SetupLogger(logger)

	if err := http_server.Start("127.0.0.1:8080"); err != nil {
		log.Fatalln()
	}
}
