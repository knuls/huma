package config

type Configurator interface {
	binder
	loader
	unmarshaller
}

type binder interface {
	SetBindings([]string)
}

type loader interface {
	Load(string, string, string) error
}

type unmarshaller interface {
	Unmarshal(interface{}) error
}
