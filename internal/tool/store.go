package tool

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/internal"
)

type ObjectType string

var (
	Commit ObjectType = "commit"
	Tree   ObjectType = "tree"
	Blob   ObjectType = "blob"
)

func StoreObject(objectType ObjectType, content string) (string, error) {
	data := fmt.Sprintf("%s %d\x00%s", objectType, len(content), content)
	fileHash, err := Hash([]byte(data))
	if err != nil {
		return "", err
	}

	dir := filepath.Join(internal.GitDir, "objects", fileHash[:2])
	file := filepath.Join(dir, fileHash[2:])

	if err = os.MkdirAll(dir, 0775); err != nil {
		return "", err
	}

	buf, err := CompressObject(data)
	if err != nil {
		return "", err
	}

	if err = os.WriteFile(file, buf, 0644); err != nil {
		return "", err
	}

	return fileHash, nil
}

func CompressObject(data string) ([]byte, error) {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	if _, err := zw.Write([]byte(data)); err != nil {
		return nil, err
	}
	defer zw.Close()
	return buf.Bytes(), nil
}
