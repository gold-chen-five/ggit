package repo

import (
	"os"
	"path/filepath"
)

var gitDir = ".ggit"

func InitRepository() error {
	paths := []string{
		gitDir,
		filepath.Join(gitDir, "objects"),
		filepath.Join(gitDir, "refs", "head"),
	}

	for _, path := range paths {
		err := os.MkdirAll(path, 0775)
		if err != nil {
			return err
		}
	}

	// Create HEAD file
	headPath := filepath.Join(gitDir, "HEAD")
	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return err
	}

	// Create config file (empty for now)
	configPath := filepath.Join(gitDir, "config")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
