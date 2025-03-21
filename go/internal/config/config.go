package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"GRPC_PORT"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT"`
}

// Loads and returns config,
// panics if fails to do so.
func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist on given path: " + configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &config
}

// Fetches config path from cli flag or
// environmental variable CONFIG_PATH.
//
// Priority: flag > env > default.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		if res = os.Getenv("CONFIG_PATH"); res == "" {
			res = "./config/dev.yml"
		}
	}

	return res
}
