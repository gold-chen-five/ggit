package repo

import (
	"os"
	"path/filepath"
	"time"
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
	//tree, err := createTreeOject()
	// if err != nil {
	// 	return err
	// }

	//parent := getParentCommitHash()
	// create commit object
	return nil
}

func createTreeOject() (string, error) {
	return "dummy-tree-hash", nil
}

func getParentCommitHash() string {
	headPath := filepath.Join(GitDir, "HEAD")
	headContent, err := os.ReadFile(headPath)
	if err != nil {
		return ""
	}

	return string(headContent)
}
