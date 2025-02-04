package repo

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gold-chen-five/ggit/core"
)

type Commit struct {
	Author   Person
	Commiter Person
	Message  string
	Tree     string
	Parent   string
}

type Person struct {
	Name  string
	Email string
	Time  time.Time
}

func CommitChanges(message string, author string) error {
	// read index file
	// indexPath := filepath.Join(internal.GitDir, "index")
	// indexContent, err := os.ReadFile(indexPath)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func createTreeOject() (string, error) {
	return "dummy-tree-hash", nil
}

func getParentCommitHash() string {
	headPath := filepath.Join(core.GitDir, "HEAD")
	headContent, err := os.ReadFile(headPath)
	if err != nil {
		return ""
	}

	return string(headContent)
}
