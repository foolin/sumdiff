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

	"github.com/foolin/sumdiff/vo"

	"github.com/foolin/sumdiff/internal/statusbar"
	"github.com/foolin/sumdiff/internal/util"
)

func HashWithArgs(args ...string) (list vo.HashList, err error) {
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

func Hash(h hash.Hash, paths ...string) (list vo.HashList, err error) {
	if len(paths) == 0 {
		return list, fmt.Errorf("at least one file is required")
	}

	//Format path
	for i, v := range paths {
		paths[i] = util.FormatPath(v)
	}

	list = make(vo.HashList, 0)
	fn := func(root, file string, size int64) error {
		relative := util.RelativePath(file, root)
		statusbar.Display("Calculating %v ...", relative)
		hex, err := util.HashHex(h, file)
		if err != nil {
			return err
		}
		ret := vo.HashInfo{
			Path: relative,
			Size: size,
			Hash: hex,
		}
		list = append(list, ret)
		return nil
	}

	for _, f := range paths {
		statusbar.Display("Preparing %v ...", f)

		root := ""

		stat, err := os.Stat(f)
		if err != nil {
			return list, err
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
				return list, err
			}
		} else {
			err = fn(root, f, stat.Size())
			if err != nil {
				return list, err
			}
		}
	}
	return list, nil
}
