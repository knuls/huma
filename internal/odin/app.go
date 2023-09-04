package odin

import (
	"github.com/knuls/huma/pkg/log"
	"github.com/knuls/huma/pkg/middleware"
	"github.com/knuls/huma/pkg/mux"
)

type app struct {
	cfg    *cfg
	logger log.Logger
}

func NewApp(cfg *cfg, logger log.Logger) *app {
	return &app{
		cfg:    cfg,
		logger: logger,
	}
}

func (a *app) Middlewares() []mux.Middleware {
	middlewares := []mux.Middleware{
		middleware.CORS(a.cfg.Security.Allowed.Origins, a.cfg.Security.Allowed.Methods, a.cfg.Security.Allowed.Headers, a.cfg.Security.AllowCredentials),
		middleware.JSON,
		middleware.RealIP,
		middleware.RequestID,
		middleware.Recoverer,
		middleware.Logger(a.logger),
	}
	return middlewares
}

func (a *app) Routes() []mux.Route {
	return nil
}
