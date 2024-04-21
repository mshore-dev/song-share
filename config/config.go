package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var ActiveConfig Config

type Config struct {
	ListenAddr string `toml:"listen_addr"`
	Database   string `toml:"database"`
}

func LoadConfig(path string) {

	_, err := toml.DecodeFile(path, &ActiveConfig)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

}
