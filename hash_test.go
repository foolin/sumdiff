package sumdiff

import (
	"crypto/sha1"
	"testing"
)

func TestHash(t *testing.T) {
	ok, err := Hash(sha1.New(), "../../test_data/")
	t.Logf("result: %v, error: %v", ok, err)
}
