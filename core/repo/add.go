package repo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/core"
	"github.com/gold-chen-five/ggit/core/tool"
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

	indexPath := filepath.Join(core.GitDir, "index")

	file, err := os.OpenFile(indexPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	entries, err := readIndexFile()
	if err != nil {
		return err
	}
	index, isFound := FindEntry(entries, path)

	newEntries := createNewEntries(isFound, fileHash, path, entries, index)

	if newEntries != nil {
		newContent := ConvertEntriesToContent(newEntries)

		if _, err = file.WriteString(newContent); err != nil {
			return err
		}
	}

	return nil
}

func createNewEntries(isFound bool, fileHash string, path string, entries []Entry, index int) []Entry {
	var newEntries []Entry
	if isFound {
		en := entries[index]
		if en.Hash != fileHash {
			entries[index] = Entry{
				Mode: 100644,
				Hash: fileHash,
				Path: path,
			}
			newEntries = entries
		}
	} else {
		newEntry := Entry{
			Mode: 100644,
			Hash: fileHash,
			Path: path,
		}
		newEntries = append(entries, newEntry) // Append new entry
	}

	return newEntries
}
