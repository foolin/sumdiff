package util

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(path string) (map[string]PathInfo, error) {
	return ListPath(path, func(info PathInfo) bool {
		if info.Info.IsDir() {
			return false
		}
		return true
	})
}

func ListPath(path string, acceptFn func(info PathInfo) bool) (map[string]PathInfo, error) {
	data := make(map[string]PathInfo)
	err := WalkPath(path, func(info PathInfo, err error) error {
		if err != nil {
			return err
		}
		if acceptFn != nil && !acceptFn(info) {
			return nil
		}
		data[info.Relative] = info
		return nil
	})
	return data, err
}

func WalkPath(root string, fn WalkFunc) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		relative := RelativePath(path, root)
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

func FormatPath(path string) string {
	if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "~\\") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			path = filepath.Join(homeDir, path[2:])
		}
	}
	return path
}

func AbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	} else {
		return absPath
	}
}

func RelativePath(path, root string) string {
	absPath := AbsPath(path)
	absRoot := AbsPath(root)
	relative := strings.TrimPrefix(absPath, absRoot)
	if len(absPath) > len(relative) {
		return strings.TrimPrefix(relative, string(filepath.Separator))
	}
	return relative
}
