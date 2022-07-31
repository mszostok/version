package upgrade

import (
	"os"
	"path/filepath"
)

// ConfigDir returns config directory.
//
// Path precedence:
// 1. VERSION_CONFIG_DIR,
// 2. XDG_CONFIG_HOME,
// 3. HOME.
func ConfigDir() string {
	if path := os.Getenv("VERSION_CONFIG_DIR"); path != "" {
		return path
	}

	if path := os.Getenv("XDG_CONFIG_HOME"); path != "" {
		return filepath.Join(path, "cli-version")
	}

	d, _ := os.UserHomeDir()
	return filepath.Join(d, ".config", "cli-version")
}
