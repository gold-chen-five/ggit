package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/internal/tool"
)

// 1. read file conetent by filepath
// 2. hash it
// 3. use first two hash characters create folder,
// 4. remain characters be the file name
func AddFileToIndex(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("fail to read file: %w", err)
	}

	// hash
	fileHash, err := tool.Hash(content)
	if err != nil {
		return err
	}

	// create folder and file
	objectDir := filepath.Join(GitDir, "objects", fileHash[:2])
	if err = os.Mkdir(objectDir, 0775); err != nil {
		return err
	}

	objectFilepath := filepath.Join(objectDir, fileHash[2:])
	if err = os.WriteFile(objectFilepath, content, 0644); err != nil {
		return err
	}

	indexPath := filepath.Join(GitDir, "index")
	indexContent := fmt.Sprintf("%s %s\n", fileHash, path)
	return os.WriteFile(indexPath, []byte(indexContent), 0644)
}
