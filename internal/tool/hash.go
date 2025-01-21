package tool

import (
	"crypto/sha1"
	"fmt"
)

func Hash(content []byte) (string, error) {
	// hash
	sha := sha1.New()
	if _, err := sha.Write(content); err != nil {
		return "", err
	}
	hashSum := sha.Sum(nil)
	fileHash := fmt.Sprintf("%x", hashSum)
	return fileHash, nil
}
