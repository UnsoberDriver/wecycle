// Package runner
package runner

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/UnsoberDriver/wecycle/internal/config"
)

func Apply(cfg *config.Config, projectDir string) error {
	bin := cfg.LinuxWallpaper.Binary
	if bin == "" {
		bin = "linux-wallpaperengine"
	}
	args := append([]string{"--directory", projectDir}, cfg.LinuxWallpaper.Args...)
	cmd := exec.Command(bin, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %w (%s)", bin, err, strings.TrimSpace(string(out)))
	}
	return nil
}
