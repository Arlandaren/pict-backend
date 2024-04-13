package main

import (
	http_server "course-backend/pkg/delivery/http"
	"log"
	"log/slog"
	"os"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	http_server := http_server.HttpServer{}
	http_server.RouteAll()
	http_server.SetupLogger(logger)

	if err := http_server.Start(os.Getenv("HTTP_ADDR")); err != nil {
		log.Fatalln()
	}
}
