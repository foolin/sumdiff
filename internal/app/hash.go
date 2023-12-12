package app

import (
	"fmt"
	"github.com/foolin/sumdiff/internal/util"
	"github.com/foolin/sumdiff/vo"
	"hash"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func Hash(h hash.Hash, paths ...string) (results []vo.HashVo, err error) {
	if len(paths) == 0 {
		return results, fmt.Errorf("at least one file is required")
	}
	results = make([]vo.HashVo, 0)
	fn := func(root, file string, size int64) error {
		relative := strings.TrimPrefix(file, root)
		hex, err := util.HashHex(h, file)
		if err != nil {
			return err
		}
		ret := vo.HashVo{
			Path: relative,
			Size: size,
			Hash: hex,
		}
		results = append(results, ret)
		return nil
	}
	for _, f := range paths {
		root := ""
		stat, err := os.Stat(f)
		if err != nil {
			return results, err
		}
		if stat.IsDir() {
			root = f
			err = filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				return fn(root, path, stat.Size())
			})
			if err != nil {
				return results, err
			}
		} else {
			err = fn(root, f, stat.Size())
			if err != nil {
				return results, err
			}
		}
	}
	return results, nil
}
