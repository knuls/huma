package odin

import "time"

var bindings = []string{
	"service.name",
	"service.port",

	"store.host",
	"store.port",
	"store.timeout",

	"server.timeout.read",
	"server.timeout.write",

	"security.allowed.origins",
	"security.allowed.methods",
	"security.allowed.headers",
	"security.allowCredentials",

	"auth.csrf",
}

type cfg struct {
	Service  serviceCfg
	Store    storeCfg
	Server   serverCfg
	Security securityCfg
	Auth     authCfg
}

type serviceCfg struct {
	Name string
	Port int
}

type storeCfg struct {
	Host    string
	Port    int
	Timeout time.Duration
}

type serverCfg struct {
	Timeout struct {
		Read  time.Duration
		Write time.Duration
	}
}

type securityCfg struct {
	Allowed struct {
		Origins []string
		Methods []string
		Headers []string
	}
	AllowCredentials bool
}

type authCfg struct {
	Csrf string
}
