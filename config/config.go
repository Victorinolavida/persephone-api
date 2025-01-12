package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")
var configPath = "config.yml"

type server struct {
	Port   int  `koanf:"port"`
	Debug  bool `koanf:"debug"`
	Pretty bool `koanf:"pretty"`
}

type Config struct {
	Server server   `koanf:"server"`
	DB     DBConfig `koanf:"database"`
}

func NewConfig() (Config, error) {
	var c Config

	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return c, err
	}

	err := k.Unmarshal("", &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
