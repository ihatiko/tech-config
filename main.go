package tech_config

import (
	"errors"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	configPath = "config/config.toml"
)

type ConfigSettings struct {
	Path string
}

type Options func(setting *ConfigSettings)

func WithPath(path string) Options {
	return func(setting *ConfigSettings) {
		setting.Path = path
	}
}

func ToConfig[T any](t T, opts ...Options) error {
	s := new(ConfigSettings)
	for _, opt := range opts {
		opt(s)
	}
	if s.Path == "" {
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		newPath, err := url.JoinPath(path, configPath)
		if err != nil {
			return errors.Join(err, errors.New("config parse function"))
		}
		s.Path = newPath
	}
	_, err := toml.DecodeFile(s.Path, t)
	return err
}
