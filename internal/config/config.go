package config

import (
	"github.com/pelletier/go-toml"
	"log"
	"time"
)

type Config struct {
	Login       string        `toml:"login"`
	Host        string        `toml:"host"`
	Port        string        `toml:"port"`
	InfoTimeout time.Duration `toml:"info_timeout"`
}

func MustLoad() *Config {
	cfg, err := toml.LoadFile("fb_cfg.toml")
	if err != nil {
		log.Fatalf("error loading config file: %s", err)
	}

	var config Config

	if err := cfg.Unmarshal(&config); err != nil {
		log.Fatalf("error decoding config: %s", err)
	}

	log.Printf("the config values: %v", config)

	return &config
}
