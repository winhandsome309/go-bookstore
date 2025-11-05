package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Config struct {
	HttpPort     string `yaml:"http_port" env:"http_port"`
	HostBackend  string `yaml:"host_backend" env:"host_backend"`
	HostFrontend string `yaml:"host_frontend" env:"host_frontend"`
	AuthSecret   string `yaml:"auth_secret" env:"auth_secret"`
	DatabaseURI  string `yaml:"database_uri" env:"database_uri"`
}

var (
	cfg Config
)

func LoadConfig() *Config {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	configPath := filepath.Join(currentDir, "config.yaml")

	if data, err := os.ReadFile(configPath); err == nil {
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			log.Errorf("Failed to parse YAML config: %v", err)
		}
	} else {
		log.Warnf("Config file not found at %s, using environment variables only", configPath)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Errorf("Failed to parse environment variables: %v", err)
	}

	if cfg.HttpPort == "" {
		log.Warn("http_port not set, defaulting to 8080")
		cfg.HttpPort = "8080"
	}
	if cfg.DatabaseURI == "" {
		log.Error("database_uri is required but not set")
	}

	return &cfg
}

func GetConfig() *Config {
	return &cfg
}
