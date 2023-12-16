package sumdiff

import (
	"testing"

	"github.com/foolin/sumdiff/internal/util"
)

func TestCmpFile(t *testing.T) {
	ok, res, err := CmpFile("../../test_data/a.txt", "../../test_data/b.txt")
	t.Logf("result: %v, result: %v, error: %v", ok, util.PettyJson(res), err)
}

func TestCmpDir(t *testing.T) {
	ok, res, err := CmpDir("../../test_data/data4", "../../test_data/data3")
	t.Logf("result: %v, result: %v, error: %v", ok, util.PettyJson(res), err)
}
