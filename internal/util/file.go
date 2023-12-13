package util

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func ListFiles(path string) (map[string]PathInfo, error) {
	return ListPath(path, true)
}

func ListPath(path string, onlyFile bool) (map[string]PathInfo, error) {
	data := make(map[string]PathInfo)
	err := WalkPath(path, func(info PathInfo, err error) error {
		if err != nil {
			return err
		}
		if onlyFile {
			return nil
		}
		data[info.Relative] = info
		return nil
	})
	return data, err
}

func WalkPath(path string, fn WalkFunc) error {
	rootPath := path
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		relative := strings.TrimPrefix(path, rootPath)
		walkInfo := PathInfo{
			Relative: relative,
			Path:     path,
			Info:     info,
		}
		return fn(walkInfo, err)
	})
}

type WalkFunc func(info PathInfo, err error) error

type PathInfo struct {
	Relative string
	Path     string
	Info     fs.FileInfo
}

func FileType(isDir bool) string {
	if isDir {
		return "dir"
	}
	return "file"
}
