package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	err := http.ListenAndServe(":80", app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
