package api

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	StagingDirectory   string   `toml:"staging-directory"`
	StorageDirectories []string `toml:"storage-directories"`
	PluginPaths        []string `toml:"plugin-paths"`
}

func LoadSettings(path string) (Settings, error) {
	settingsBytes, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, err
	}
	var settings Settings
	if err := toml.Unmarshal(settingsBytes, &settings); err != nil {
		return Settings{}, err
	}
	return settings, nil
}
