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
	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}
