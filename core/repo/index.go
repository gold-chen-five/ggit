package repo

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) ([]Entry, error) {
	// indexPath := filepath.Join(core.GitDir, "index")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	entries, err := ToEntries(lines)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
