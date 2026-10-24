// Package workshop
package workshop

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
)

func FindDirs() []string {
	home, _ := os.UserHomeDir()
	var dirs []string

	candidates := []string{
		filepath.Join(home, ".steam", "steam", "steamapps", "workshop", "content", "431960"),
		filepath.Join(home, ".local", "share", "Steam", "steamapps", "workshop", "content", "431960"),
		filepath.Join(home, ".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam", "steamapps", "workshop", "content", "431960"),
		filepath.Join(xdg.DataHome, "Steam", "steamapps", "workshop", "content", "431960"),
	}

	for _, c := range candidates {
		if _, err := os.Stat(c); err == nil {
			dirs = append(dirs, c)
		}
	}
	return dirs
}

func Scan(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		// skip obvious non-wallpaper dirs
		if strings.HasPrefix(name, ".") {
			continue
		}
		out = append(out, filepath.Join(dir, name))
	}
	return out, nil
}
