package mongo

type Config struct {
	Db      string `yaml:"db"`
	Uri     string `yaml:"uri"`
	Timeout string `yaml:"timeout,omitempty"`
	MinPool string `yaml:"min-pool,omitempty"`
	MaxPool string `yaml:"max-pool,omitempty"`
}
