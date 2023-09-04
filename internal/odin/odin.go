package odin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/knuls/huma/pkg/config"
	"github.com/knuls/huma/pkg/log"
	"github.com/knuls/huma/pkg/mux"
	"github.com/knuls/huma/pkg/store"
)

func New() {
	logger := log.NewStructuredLogger()

	cfger := config.NewViperConfig("odin")
	cfger.SetBindings(bindings)
	var cfg *cfg
	if err := cfger.LoadAndUnmarshal("config.odin", "yaml", filepath.Join("configs"), &cfg); err != nil {
		logger.Error("config load and unmarhsal", "error", err)
		return
	}

	store, err := store.NewMongoStore(cfg.Store.Host, cfg.Store.Port, cfg.Store.Timeout)
	if err != nil {
		logger.Error("store new", "error", err)
		return
	}
	defer func() {
		if err := store.Disconnect(); err != nil {
			logger.Error("store disconnect", "error", err)
			return
		}
	}()

	app := NewApp(cfg, logger)

	mux := mux.NewChiMux()
	mux.Middlewares(app.Middlewares()...)
	mux.Routes(app.Routes()...)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Service.Port),
		Handler:      mux.Handler(),
		ErrorLog:     logger.GetStdLogger(),
		ReadTimeout:  cfg.Server.Timeout.Read,
		WriteTimeout: cfg.Server.Timeout.Write,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("server listen", "error", err)
			return
		}
	}()

	logger.Info(fmt.Sprintf("starting %s service on port %d", "", 0))

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh

	logger.Info(fmt.Sprintf("received shutdown signal %s", sig.String()))

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error("server shutdown", "error", err)
		return
	}
}
