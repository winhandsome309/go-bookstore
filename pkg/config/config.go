package config

import (
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	HttpPort     string `env:"http_port"`
	HostBackend  string `env:"host_backend"`
	HostFrontend string `env:"host_frontend"`
	AuthSecret   string `env:"auth_secret"`
	DatabaseURI  string `env:"database_uri"`
}

var (
	cfg Config
)

func LoadConfig() *Config {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil {
		log.Error(err)
	}
	if err := env.Parse(&cfg); err != nil {
		log.Error(err)
	}
	return &cfg
}

func GetConfig() *Config {
	return &cfg
}
