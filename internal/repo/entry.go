package repo

import (
	"sort"
	"strings"
)

type Entry struct {
	Mode string
	Hash string
	Path string
}

func ToEntries(lines []string) []Entry {
	var indexEntries []Entry
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		var entry Entry
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
