package vo

import "fmt"

type HashVo struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	Hash string `json:"hash"`
}

func HashToTable(list []HashVo) [][]string {
	out := make([][]string, len(list))
	for i, v := range list {
		out[i] = []string{v.Hash, fmt.Sprintf("%v", v.Size), v.Path}
	}
	return out
}
