package jsonlog

import (
	"context"
	"io"
	"log/slog"
	"runtime/debug"
)

type Logger struct {
	level   slog.Level
	log     *slog.Logger
	handler slog.Handler
}

func New(out io.Writer, level slog.Level) *Logger {
	handler := slog.NewJSONHandler(out, &slog.HandlerOptions{
		Level: level,
	})

	return &Logger{
		level:   level,
		log:     slog.New(handler),
		handler: handler,
	}
}

func (l *Logger) Handler() slog.Handler {
	return l.handler
}

func (l *Logger) Log(msg string, properties map[string]any) {
	l.print(msg, properties)
}

func (l *Logger) Error(err error, properties map[string]any) {
	l.print(err.Error(), properties)
}

func (l *Logger) print(msg string, props map[string]any) {
	aux := struct {
		Properties map[string]any `json:"properties,omitempty"`
		Trace      string         `json:"trace,omitempty"`
	}{
		Properties: props,
	}

	if l.level >= slog.LevelError {
		aux.Trace = string(debug.Stack())
	}

	l.log.Log(context.Background(), l.level, msg, slog.Any("props", aux))
}
