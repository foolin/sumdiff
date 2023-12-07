package app

import (
	"github.com/foolin/sumdiff/internal/util"
	"io/fs"
	"os"
)

func DiffDirectory(path1, path2 string) Result {
	util.WalkDir(path1, func(relative string, path string, info fs.FileInfo, err error) error {

	})

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
	h1, err := util.Sha1(file1)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	h2, err := util.Sha1(file2)
	if err != nil {
		return Result{CODE_ERROR, err}
	}
	if h1 != h2 {
		return Result{CODE_NQ_HASH, err}
	}
	return OK
}
