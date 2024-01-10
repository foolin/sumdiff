package vo

import "fmt"

type HashInfo struct {
	Path  string `json:"path"`
	Size  int64  `json:"size"`
	Hash  string `json:"hash"`
	Error error  `json:"error,omitempty" yaml:"error,omitempty"`
}

type HashList []HashInfo

func (r HashList) Array() [][]string {
	out := make([][]string, len(r)+1)
	out[0] = []string{"HASH", "SIZE", "PATH"}
	for i, v := range r {
		out[i+1] = []string{v.Hash, fmt.Sprintf("%v", v.Size), v.Path}
	}
	return out
}
