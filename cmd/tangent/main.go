// Command tangent runs the platform API, the provisioning worker, or database migrations.
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

	"github.com/thetangentline/tangent/internal/config"
)

const usage = "usage: tangent api | worker | migrate"

// version is set at build time via -ldflags.
var version = "dev"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	mode := config.Mode(os.Args[1])

	cfg, err := config.Load(mode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "config: %v\n", err)
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: cfg.LogLevel}))
	slog.SetDefault(logger.With("mode", mode, "version", version))

	if err := run(mode, cfg); err != nil {
		slog.Error("exited with error", "err", err)
		os.Exit(1)
	}
}

func run(mode config.Mode, cfg *config.Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	switch mode {
	case config.ModeAPI:
		return runAPI(ctx, cfg)
	case config.ModeWorker:
		return runWorker(ctx, cfg)
	case config.ModeMigrate:
		return runMigrate(ctx, cfg)
	}
	return fmt.Errorf("unknown mode %q", mode)
}

func runAPI(ctx context.Context, cfg *config.Config) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		slog.Info("api listening", "addr", cfg.HTTPAddr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	slog.Info("api shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}

func runWorker(ctx context.Context, cfg *config.Config) error {
	// TODO(M3): provisioning runner.
	slog.Info("worker started", "concurrency", cfg.WorkerConcurrency)
	<-ctx.Done()
	slog.Info("worker shutting down")
	return nil
}

func runMigrate(_ context.Context, _ *config.Config) error {
	// TODO(M2): apply migrations/ to the database.
	slog.Info("no migrations yet")
	return nil
}
