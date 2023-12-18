package sumdiff

import (
	"testing"
)

func TestDiffFile(t *testing.T) {
	ok, err := EqualFile("../../test_data/a.txt", "../../test_data/b.txt")
	t.Logf("result: %v, error: %v", ok, err)
}

func TestDiffDir(t *testing.T) {
	ok, err := EqualDir("../../test_data/data4", "../../test_data/data3")
	t.Logf("result: %v, error: %v", ok, err)
}
