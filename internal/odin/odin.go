package odin

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/knuls/huma/pkg/core/config"
	"github.com/knuls/huma/pkg/core/logger"
	"github.com/knuls/huma/pkg/core/middleware"
	"github.com/knuls/huma/pkg/core/validator"
	"github.com/knuls/huma/pkg/creator"
	"github.com/knuls/huma/pkg/organization"
	"github.com/knuls/huma/pkg/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New() {
	// logger
	log, err := logger.NewZapLogger()
	if err != nil {
		fmt.Printf("logger new error: %v", err)
		os.Exit(1)
	}
	defer log.GetZapLogger().Sync()

	// config
	cfg := config.New("odin")
	cfg.SetBindings(bindings)
	var appCfg *appConfig
	if err := cfg.Load("config.odin", "yaml", filepath.Join("configs"), &appCfg); err != nil {
		log.Error("config load", "error", err)
		return
	}

	// db
	dbCtx, cancel := context.WithTimeout(context.Background(), appCfg.Store.Timeout*time.Second)
	defer cancel()
	uri := fmt.Sprintf("%s://%s:%d", appCfg.Store.Client, appCfg.Store.Host, appCfg.Store.Port)
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Error("db connect", "error", err)
		return
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Error("db disconnect", "error", err)
			return
		}
	}()
	pingCtx, cancel := context.WithTimeout(context.Background(), appCfg.Store.Timeout*time.Second)
	defer cancel()
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Error("db ping", "error", err)
		return
	}

	// mux
	mux := chi.NewRouter()

	// middlewares
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   appCfg.Security.Allowed.Origins,
		AllowedMethods:   appCfg.Security.Allowed.Methods,
		AllowedHeaders:   appCfg.Security.Allowed.Headers,
		AllowCredentials: appCfg.Security.AllowCredentials,
	}))
	mux.Use(middleware.JSON)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger(log))

	// validator
	val, err := validator.New()
	if err != nil {
		log.Error("validator new", "error", err)
		return
	}

	// dao's
	creatorDao := creator.NewDao(client, val)
	organizationDao := organization.NewDao(client, val)
	schemaDao := schema.NewDao(client, val)

	// svc's
	creatorSvc := creator.NewService(creatorDao)
	organizationSvc := organization.NewService(organizationDao)
	schemaSvc := schema.NewService(schemaDao)

	// routers
	mux.Mount("/creator", creator.NewMux(creatorSvc).Routes())
	mux.Mount("/organization", organization.NewMux(organizationSvc).Routes())
	mux.Mount("/schema", schema.NewMux(schemaSvc).Routes())

	// server
	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", appCfg.Service.Port),
		Handler:      mux,
		ErrorLog:     log.GetStdLogger(),
		ReadTimeout:  appCfg.Server.Timeout.Read * time.Second,
		WriteTimeout: appCfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  appCfg.Server.Timeout.Idle * time.Second,
	}

	// listen
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("server listen", "error", err)
			return
		}
	}()
	log.Infof("starting %s service on port: %d", appCfg.Service.Name, appCfg.Service.Port)

	// shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh
	fmt.Println(sig.String())

	shutdownCtx, cancel := context.WithTimeout(context.Background(), appCfg.Server.Timeout.Shutdown*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return
	}
}
