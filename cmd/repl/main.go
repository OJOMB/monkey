package main

import (
	"log/slog"
	"os"

	"github.com/OJOMB/donkey/internal/repl"
	"github.com/OJOMB/donkey/pkg/logs"
)

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	r := repl.New(os.Stdin, os.Stdout, logs.NewMultiSlogger(jsonHandler))
	r.Start()
}
