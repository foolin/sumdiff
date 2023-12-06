package app

import (
	"os"

	"github.com/foolin/sumdiff/internal/utilx"
)

func Diff(path1, path2 string) bool {

}

func DiffFile(file1, file2 string) Result {
	f1, err := os.Stat(file1)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	f2, err := os.Stat(file2)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	if f1.Size() != f2.Size() {
		return Result{CODE_NQ_SIZE, err}
	}
	h1, err := utilx.Sha1(file1)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	h2, err := utilx.Sha1(file2)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	if h1 != h2 {
		return Result{CODE_NQ_HASH, err}
	}
	return OK
}
