package main

import (
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"zikr-app/internal/zikr/port"
	"zikr-app/pkg/config"
	db "zikr-app/pkg/db"
	"zikr-app/pkg/logger/slogpretty"
)

const (
	envLocal = "local"
	envProd  = "prod"
	//envDev   = "dev"
)

func main() {
	// Config
	cfg := config.Load()

	// Logger
	log := setupLogger("local")

	// DB
	pgxConn, err := db.ConnDB()
	if err != nil {
		log.Error("failed to connect to database", err)
		os.Exit(1)
	}
	defer pgxConn.Close()

	router := port.New(port.RouterOption{
		DB: pgxConn,
	})

	log.Info("starting server", slog.String("address", cfg.HttpPort))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + cfg.HttpPort,
		Handler: router,
		//ReadTimeout:  cfg.HTTPServer.Timeout,
		//WriteTimeout: cfg.HTTPServer.Timeout,
		//IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("failed to start server", err.Error())
		}
	}()

	log.Info("server started")

	<-done
	log.Info("stopping server")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
