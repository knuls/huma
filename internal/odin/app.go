package odin

import (
	"github.com/knuls/huma/pkg/creator"
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
		middleware.AllowContentType("application/json"),
		middleware.JSON,
		middleware.NoCache,
		middleware.RealIP,
		middleware.RequestID,
		middleware.Logger(a.logger),
		middleware.Recoverer,
	}
	return middlewares
}

func (a *app) Routes() []mux.Route {
	return []mux.Route{
		{Pattern: "/creator", Handler: creator.NewMux().Routes()},
	}
}
