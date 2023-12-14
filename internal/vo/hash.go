package vo

import "fmt"

type HashVo struct {
	Path  string `json:"path"`
	Size  int64  `json:"size"`
	Hash  string `json:"hash"`
	Error error  `json:"error"`
}

func HashToTable(list []HashVo) [][]string {
	out := make([][]string, len(list)+1)
	out[0] = []string{"Hash", "Size", "Path"}
	for i, v := range list {
		out[i+1] = []string{v.Hash, fmt.Sprintf("%v", v.Size), v.Path}
	}
	return out
}
