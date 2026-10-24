// Package config
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

type Config struct {
	Path              string            `toml:"-"`
	Settings          Settings          `toml:"settings"`
	LinuxWallpaper    LinuxWallpaper    `toml:"linux_wallpaperengine"`
}

type Settings struct {
	Mode       string   `toml:"mode"`
	WorkshopDir string  `toml:"workshop_dir"`
	Exclude    []string `toml:"exclude"`
}

type LinuxWallpaper struct {
	Binary string   `toml:"binary"`
	Args   []string `toml:"args"`
}

func DefaultPath() string {
	return filepath.Join(xdg.ConfigHome, "wecycle", "config.toml")
}

func ExpandPath(p string) string {
	if p == "" {
		return p
	}
	if p == "~" || strings.HasPrefix(p, "~/") {
		home, _ := os.UserHomeDir()
		if p == "~" {
			return home
		}
		return filepath.Join(home, strings.TrimPrefix(p, "~/"))
	}
	return p
}

func Load(path string) (*Config, error) {
	if path == "" {
		path = DefaultPath()
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg := &Config{}
			cfg.applyDefaults()
			return cfg, nil
		}
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg Config
	if _, err := toml.Decode(string(raw), &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	cfg.Path = path
	cfg.applyDefaults()
	return &cfg, nil
}

func (c *Config) applyDefaults() {
	if c.Settings.Mode == "" {
		c.Settings.Mode = "random"
	}
	if c.LinuxWallpaper.Binary == "" {
		c.LinuxWallpaper.Binary = "linux-wallpaperengine"
	}
	if c.Settings.WorkshopDir != "" {
		c.Settings.WorkshopDir = ExpandPath(c.Settings.WorkshopDir)
	}
}
