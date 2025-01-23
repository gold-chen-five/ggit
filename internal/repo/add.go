package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/internal"
	"github.com/gold-chen-five/ggit/internal/tool"
)

// 1. read file conetent by filepath
// 2. store file content as object
func AddFileToIndex(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("fail to read file: %w", err)
	}

	fileHash, err := tool.StoreObject(content)
	if err != nil {
		return err
	}

	indexPath := filepath.Join(internal.GitDir, "index")
	indexContent := fmt.Sprintf("%d %s %s\n", 100644, fileHash, path)

	// write to index file
	file, err := os.OpenFile(indexPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write([]byte(indexContent)); err != nil {
		return err
	}

	return nil
}
