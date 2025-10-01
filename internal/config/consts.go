package config

import (
	"path"

	"github.com/adrg/xdg"
)

var DefaultConfigFile string = path.Join(xdg.ConfigHome, "gelo", "config.json")
