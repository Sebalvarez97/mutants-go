package db

type Config struct {
	Type    Type   `yaml:type`
	Name    string `yaml:"name"`
	Uri     string `yaml:"uri"`
	Timeout string `yaml:"timeout,omitempty"`
	MinPool string `yaml:"min-pool,omitempty"`
	MaxPool string `yaml:"max-pool,omitempty"`
}

type Type string

const (
	MongoDB = "mongo"
	Mysql   = "mysql"
)
