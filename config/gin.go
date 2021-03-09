package config

type Gin struct {
	BaseRouter Router `yaml:"router"`
}

type Router struct {
	Port     string `yaml:"port"`
	BasePath string `yaml:"base-path"`
}
