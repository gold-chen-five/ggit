package repo

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/gold-chen-five/ggit/internal"
)

type IndexEntry struct {
	Mode string
	Hash string
	Path string
}

func readIndexFile() ([]IndexEntry, error) {
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

	return transferToIndexEntry(lines), nil
}

func transferToIndexEntry(lines []string) []IndexEntry {
	var indexEntries []IndexEntry
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		var entry IndexEntry
		for index, s := range splitLine {
			switch index {
			case 0:
				entry.Mode = s
			case 1:
				entry.Hash = s
			case 2:
				entry.Path = s
			}
		}
		indexEntries = append(indexEntries, entry)
	}
	return indexEntries
}
