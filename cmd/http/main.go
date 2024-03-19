package main

import (
	"flag"
	"github.com/maks-herasymov/solidgate/internal/env"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"
)

// @title Solidgate Test Task
// @version 1.0
// @description Some api for this test task.

// @BasePath /
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL  string
	httpPort int
}

type application struct {
	config config
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:8080")
	cfg.httpPort = env.GetInt("HTTP_PORT", 8080)

	flag.Parse()

	app := &application{
		config: cfg,
		logger: logger,
	}

	return app.serveHTTP()
}
