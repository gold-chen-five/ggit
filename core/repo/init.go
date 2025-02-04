package repo

import (
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/core"
)

func InitRepository() error {
	paths := []string{
		core.GitDir,
		filepath.Join(core.GitDir, "objects"),
		filepath.Join(core.GitDir, "refs", "head"),
	}

	for _, path := range paths {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			return err
		}
	}

	// Create HEAD file
	headPath := filepath.Join(core.GitDir, "HEAD")
	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return err
	}

	// Create config file (empty for now)
	configPath := filepath.Join(core.GitDir, "config")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
