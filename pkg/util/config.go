package util

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config has all the config (aside from credentials) for what to do.
type Config struct {
	Albums     []*ConfigAlbum   `yaml:"albums"`
	Every      string           `yaml:"every,omitempty"`
	Prometheus ConfigPrometheus `yaml:"prometheus,omitempty"`
	Pprof      ConfigPprof      `yaml:"pprof,omitempty"`
}

type ConfigAlbum struct {
	Name         string             `yaml:"name"`
	MaxPlaylist  int            	`yaml:"maxPlaylist"`
	DryRun       *bool              `yaml:"dryRun,omitempty"`
	Delete       *bool              `yaml:"delete,omitempty"`
	ForcePublish *bool              `yaml:"forcePublish,omitempty"`
	Sources      ConfigAlbumSources `yaml:"sources"`
}

type ConfigAlbumSources struct {
	Googlephotos []string `json:"googlephotos,omitempty"`
}

type ConfigPrometheus struct {
	Listen string `yaml:"listen"`
}

type ConfigPprof struct {
	Listen string `yaml:"listen"`
}

func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := Config{}
	if err = yaml.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
