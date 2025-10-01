package config

import (
	"path"

	"github.com/adrg/xdg"
)

var (
	DefaultConfigHome string = path.Join(xdg.ConfigHome, "gelo")
	DefaultConfigFile string = path.Join(DefaultConfigHome, "config.json")
)
