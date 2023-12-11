package app

import (
	"fmt"
	"github.com/foolin/sumdiff/internal/util"
	"github.com/foolin/sumdiff/vo"
	"hash"
	"io/fs"
	"os"
	"path/filepath"
)

func Hash(h hash.Hash, paths ...string) (results []vo.HashVo, err error) {
	if len(paths) == 0 {
		return results, fmt.Errorf("at least one file is required")
	}
	results = make([]vo.HashVo, 0)
	fn := func(f string, size int64) error {
		h, err := util.HashHex(h, f)
		if err != nil {
			return err
		}
		ret := vo.HashVo{
			Path: f,
			Size: size,
			Hash: h,
		}
		results = append(results, ret)
		return nil
	}
	for _, f := range paths {
		stat, err := os.Stat(f)
		if err != nil {
			return results, err
		}
		if stat.IsDir() {
			err = filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				return fn(f, stat.Size())
			})
			if err != nil {
				return results, err
			}
		} else {
			err = fn(f, stat.Size())
			if err != nil {
				return results, err
			}
		}
	}
	return results, nil
}
