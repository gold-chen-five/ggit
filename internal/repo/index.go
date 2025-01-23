package repo

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/internal"
)

type IndexEntry struct {
	Mode string
	Hash string
	Path string
}

func readIndexFile() ([]string, error) {
	indexPath := filepath.Join(internal.GitDir, "index")
	file, err := os.Open(indexPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func transferToIndexEntry(lines []string) []IndexEntry {
	var indexEntries []IndexEntry
	for _, line := range lines {

	}
}
