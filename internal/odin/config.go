package odin

import "time"

// bindings is an array of keys used to generate
// the environment variable names when loading the configuration.
// For example: "service.name" will be -> "SERVICE_NAME"
var bindings = []string{
	"service.name",
	"service.port",

	"store.client",
	"store.host",
	"store.port",
	"store.timeout",

	"server.timeout.read",
	"server.timeout.write",
	"server.timeout.idle",
	"server.timeout.shutdown",

	"security.allowed.origins",
	"security.allowed.methods",
	"security.allowed.headers",
	"security.allowCredentials",

	"auth.csrf",
}

type appConfig struct {
	Service  serviceConfig
	Store    storeConfig
	Server   serverConfig
	Security securityConfig
	Auth     authConfig
}

type serviceConfig struct {
	Name string
	Port int
}

type storeConfig struct {
	Client  string
	Host    string
	Port    int
	Timeout time.Duration
}

type serverConfig struct {
	Timeout struct {
		Read     time.Duration
		Write    time.Duration
		Idle     time.Duration
		Shutdown time.Duration
	}
}

type securityConfig struct {
	Allowed struct {
		Origins []string
		Methods []string
		Headers []string
	}
	AllowCredentials bool
}

type authConfig struct {
	Csrf string
}
