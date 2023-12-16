package sumdiff

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/foolin/sumdiff/internal/util"
	"github.com/foolin/sumdiff/internal/vo"
)

func HashWithArgs(args ...string) (results []vo.HashVo, err error) {
	t := strings.ToLower(args[0])
	var algo hash.Hash
	switch t {
	case "md5":
		algo = md5.New()
	case "sha1":
		algo = sha1.New()
	case "sha256":
		algo = sha256.New()
	case "sha512":
		algo = sha512.New()
	default:
		return nil, fmt.Errorf("not support algo=%v", t)
	}
	return Hash(algo, args[1:]...)
}

func Hash(h hash.Hash, paths ...string) (results []vo.HashVo, err error) {
	if len(paths) == 0 {
		return results, fmt.Errorf("at least one file is required")
	}
	results = make([]vo.HashVo, 0)
	fn := func(root, file string, size int64) error {
		relative := strings.TrimPrefix(strings.TrimPrefix(file, root), string(os.PathSeparator))
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
			root = filepath.Dir(f)
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
