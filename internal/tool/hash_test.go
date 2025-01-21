package tool

import "testing"

func TestHash(t *testing.T) {
	testContent := []byte("test content")
	expectedHash := "1eebdf4fdc9fc7bf283031b93f9aef3338de9052"

	hash, err := Hash(testContent)
	if err != nil {
		t.Fatalf("Hash failed: %v", err)
	}

	if hash != expectedHash {
		t.Fatalf("Hash does not match. Expected: %s, Got: %s", expectedHash, hash)
	}
}
