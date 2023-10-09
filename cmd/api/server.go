package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		s := <-quit

		app.logger.Log("caught signal", map[string]any{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		app.logger.Log("completing background tasks", map[string]any{
			"addr": srv.Addr,
		})

		app.wg.Wait()
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.logger.Log("starting server", map[string]any{
		"addr": srv.Addr,
		"env":  app.cfg.env,
	})

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err := <-shutdownError
	if err != nil {
		return err
	}

	app.logger.Log("stopped server", map[string]any{
		"addr": srv.Addr,
	})

	return nil
}
