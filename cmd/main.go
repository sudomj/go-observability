package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	configureLogs()
	port := os.Getenv("APP_PORT")

	server := http.NewServeMux()
	server.HandleFunc("/api/v1/events", getEvents)

	slog.Info(fmt.Sprintf("Server will start at port %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), server)
}

func configureLogs() {

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func getEvents(w http.ResponseWriter, r *http.Request) {

	slog.Info("Request received",
		"method", r.Method,
		"path", r.URL.Path,
	)

	w.WriteHeader(http.StatusOK)
}
