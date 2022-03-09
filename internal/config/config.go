package config

import (
	"log"

	"github.com/amirhnajafiz/Notifier/internal/cache"
	"github.com/amirhnajafiz/Notifier/internal/client"
	"github.com/amirhnajafiz/Notifier/internal/cmd/server"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	Client client.Config `koanf:"mqtt"`
	Server server.Config `koanf:"serv"`
	Cache  cache.Config  `koanf:"cache"`
}

func New() Config {
	var instance Config

	k := koanf.New(".")

	// Load default configurations
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading defaults: %s\n", err)
	}

	// Load configuration
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s\n", err)
	}

	// Unmarshalling configurations
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s\n", err)
	}

	return instance
}
