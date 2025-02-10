package repo

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Mode uint32
	Hash string
	Path string
}

func ToEntries(lines []string) ([]Entry, error) {
	var indexEntries []Entry
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		var entry Entry
		for index, s := range splitLine {
			switch index {
			case 0:
				suint, err := strconv.ParseUint(s, 10, 32)
				if err != nil {
					return nil, err
				}
				entry.Mode = uint32(suint)
			case 1:
				entry.Hash = s
			case 2:
				entry.Path = s
			}
		}
		indexEntries = append(indexEntries, entry)
	}
	return indexEntries, nil
}

// use binary search to find entry by path, return (index, isFound)
func FindEntry(entries []Entry, path string) (index int, isFound bool) {
	index = sort.Search(len(entries), func(i int) bool {
		return entries[i].Path >= path
	})

	if index < len(entries) && entries[index].Path == path {
		return index, true
	}

	return index, false
}

func ConvertEntriesToContent(entries []Entry) string {
	var content string
	for _, newEntry := range entries {
		enStr := fmt.Sprintf("%d %s %s\n", 100644, newEntry.Hash, newEntry.Path)
		content += enStr
	}
	return content
}
