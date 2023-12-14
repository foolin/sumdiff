package app

import (
	"github.com/foolin/sumdiff/internal/util"
	"testing"
)

func TestCmpFile(t *testing.T) {
	ok, res, err := CmpFile("../../test_data/a.txt", "../../test_data/b.txt")
	t.Logf("result: %v, result: %v, error: %v", ok, util.PettyJson(res), err)
}

func TestCmpDir(t *testing.T) {
	ok, res, err := CmpDir("../../test_data/data4", "../../test_data/data3")
	t.Logf("result: %v, result: %v, error: %v", ok, util.PettyJson(res), err)
}
