package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	configuration Config
	once          sync.Once
)

const (
	envGoEnvironment = "GO_ENVIRONMENT"
	envScope         = "SCOPE"
	envPort          = "PORT"

	filePathFormat = "%s/config/yml/%s.yml"
	localFileName  = "local"
)

type Config struct {
	//Config attributes
	GinConfig Gin `yaml:"gin"`
}

func GetConfig() Config {
	once.Do(func() {

		fname := getFileName()

		configFile, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalf("yamlFile.Get err   #%v ", err)
		}

		configuration = Config{}
		if err := yaml.Unmarshal(configFile, &configuration); err != nil {
			log.Fatalf("error: %v", err)
		}

	})

	return configuration
}

func getFileName() string {
	env, scope, _ := readEnv()

	//Extract the config name from scope
	fname := strings.Split(scope, "-")[0]

	var basePath string
	if env != "production" {
		fname = localFileName
		_, file, _, _ := runtime.Caller(1)
		basePath = strings.TrimSuffix(file, "/config/config.go")
	} else {
		basePath, _ = os.Getwd()
	}

	filePath := fmt.Sprintf(filePathFormat, basePath, fname)

	return filePath
}

func readEnv() (env, scope, port string) {
	env = os.Getenv(envGoEnvironment)
	scope = os.Getenv(envScope)
	port = os.Getenv(envPort)

	if port == "" {
		port = "8080"
	}
	return
}
