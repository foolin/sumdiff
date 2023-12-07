package util

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type FileSum struct {
	Info fs.FileInfo
	Hash string
}

func WalkDir(path string, fn WalkFunc) error {
	dirName := filepath.Dir(path)
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		relative := strings.TrimPrefix(path, dirName)
		return fn(relative, path, info, err)
	})
}

type WalkFunc func(relative string, path string, info fs.FileInfo, err error) error
