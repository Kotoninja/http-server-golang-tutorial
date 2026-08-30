package main

import (
	"flag"
	"log/slog"

	"github.com/BurntSushi/toml"
	"github.com/kotoninja/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		slog.Error(err.Error())
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		slog.Error(err.Error())
	}
}
