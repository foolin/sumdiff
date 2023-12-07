package app

import "testing"

func TestDiffFile(t *testing.T) {
	result := DiffFile("../../test_data/a.txt", "../../test_data/b.txt")
	t.Logf("result: %v", result)
}
