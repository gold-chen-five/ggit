package tool

import (
	"bytes"
	"compress/zlib"
	"os"
	"path/filepath"

	"github.com/gold-chen-five/ggit/core"
)

type ObjectType string

var (
	Commit ObjectType = "commit"
	Tree   ObjectType = "tree"
	Blob   ObjectType = "blob"
)

// 1. hash it
// 2. use first two hash characters create folder,
// 3. remain characters be the file name
// 4. compress content
// 5. store the content in the file objects/xx/xxxxxx
func StoreObject(data []byte) (string, error) {
	fileHash, err := Hash(data)
	if err != nil {
		return "", err
	}

	dir := filepath.Join(core.GitDir, "objects", fileHash[:2])
	file := filepath.Join(dir, fileHash[2:])

	if err = os.MkdirAll(dir, 0775); err != nil {
		return "", err
	}

	buf, err := compressObject(data)
	if err != nil {
		return "", err
	}

	if err = os.WriteFile(file, buf, 0644); err != nil {
		return "", err
	}

	return fileHash, nil
}

func compressObject(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := zlib.NewWriter(&buf)
	if _, err := zw.Write(data); err != nil {
		return nil, err
	}
	zw.Close()
	return buf.Bytes(), nil
}
