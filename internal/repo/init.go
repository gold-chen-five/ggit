package repo

import (
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/internal"
)

func InitRepository() error {
	paths := []string{
		internal.GitDir,
		filepath.Join(internal.GitDir, "objects"),
		filepath.Join(internal.GitDir, "refs", "head"),
	}

	for _, path := range paths {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			return err
		}
	}

	// Create HEAD file
	headPath := filepath.Join(internal.GitDir, "HEAD")
	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return err
	}

	// Create config file (empty for now)
	configPath := filepath.Join(internal.GitDir, "config")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
