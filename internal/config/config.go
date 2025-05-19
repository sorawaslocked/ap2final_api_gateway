package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/pkg/grpc"
	"os"
)

type (
	Config struct {
		HTTPServer HTTPServer  `yaml:"httpServer" env-required:"true"`
		GRPC       grpc.Config `yaml:"grpc" env-required:"true"`
	}

	HTTPServer struct {
		Address      string `yaml:"address" env:"HTTP_ADDRESS" env-required:"true"`
		ReadTimeout  string `yaml:"readTimeout" env:"HTTP_READ_TIMEOUT" env-default:"30s"`
		WriteTimeout string `yaml:"writeTimeout" env:"HTTP_WRITE_TIMEOUT" env-default:"30s"`
		IdleTimeout  string `yaml:"idleTimeout" env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
		GinMode      string `yaml:"ginMode" env:"GIN_MODE" env-default:"debug"`
	}
)

func MustLoad() *Config {
	cfgPath := fetchConfigPath()

	if cfgPath == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic("config file does not exist: " + cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic("failed to load config")
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
