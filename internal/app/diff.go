package app

import (
	"fmt"
	"github.com/foolin/sumdiff/internal/plog"
	"github.com/foolin/sumdiff/internal/util"
	"os"
	"strings"
)

func Diff(path1, path2 string) (bool, error) {
	s1, err := os.Stat(path1)
	if err != nil {
		return false, err
	}
	s2, err := os.Stat(path2)
	if err != nil {
		return false, err
	}
	if s1.IsDir() != s2.IsDir() {
		return false, fmt.Errorf("not equal is directory [%v!=%v]", s1.IsDir(), s2.IsDir())
	}
	if s1.IsDir() {
		return DiffDir(path1, path2)
	} else {
		return DiffFile(path1, path2)
	}
}

func DiffDir(path1, path2 string) (bool, error) {
	data1, err := util.ListPath(path1)
	if err != nil {
		return false, err
	}
	data2, err := util.ListPath(path2)
	if err != nil {
		return false, err
	}
	if len(data1) != len(data2) {
		return false, fmt.Errorf("not equal file count [%v!=%v]", len(data1), len(data2))
	}
	for k, v1 := range data1 {
		plog.Progress("compare diff path " + k)
		v2, ok := data2[k]
		if !ok {
			return false, fmt.Errorf("path2 not exist path %v", k)
		}
		if v1.Info.Size() != v2.Info.Size() {
			return false, fmt.Errorf("%v not equal size [%v!=%v]", k, v1.Info.Size(), v2.Info.Size())
		}
		if v1.Info.IsDir() != v2.Info.IsDir() {
			return false, fmt.Errorf("%v not equal is directory [%v!=%v]", k, v1.Info.IsDir(), v2.Info.IsDir())
		}
		//File check file md5
		if !v1.Info.IsDir() {
			h1, err := util.Sha256(v1.Path)
			if err != nil {
				return false, err
			}
			h2, err := util.Sha256(v2.Path)
			if err != nil {
				return false, err
			}
			if h1 != h2 {
				return false, fmt.Errorf("%v not equal hash [%v!=%v]", k, h1, h2)
			}
		}

		delete(data2, k)
	}
	if len(data2) != 0 {
		remain := make([]string, 0)
		for _, info := range data2 {
			if len(remain) >= 5 {
				remain = append(remain, "...")
				break
			}
			remain = append(remain, info.Path)
		}
		return false, fmt.Errorf("path1 not exist path: %v", strings.Join(remain, ","))
	}
	return true, nil
}

func DiffFile(file1, file2 string) (bool, error) {
	f1, err := os.Stat(file1)
	if err != nil {
		return false, err
	}
	f2, err := os.Stat(file2)
	if err != nil {
		return false, err
	}
	if f1.Size() != f2.Size() {
		return false, fmt.Errorf("not equal size")
	}
	h1, err := util.Sha256(file1)
	if err != nil {
		return false, err
	}
	h2, err := util.Sha256(file2)
	if err != nil {
		return false, err
	}
	if h1 != h2 {
		return false, fmt.Errorf("not equal hash")
	}
	return true, nil
}
