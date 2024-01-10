package vo

import "fmt"

type CmpInfo struct {
	Hash1 *HashInfo `json:"hash1"`
	Hash2 *HashInfo `json:"hash2"`
	Equal bool      `json:"equal"`
	Msg   string    `json:"msg"`
}

func NewCmpInfo(path1, path2 string) CmpInfo {
	return CmpInfo{
		Hash1: &HashInfo{
			Path:  path1,
			Size:  0,
			Hash:  "",
			Error: nil,
		},
		Hash2: &HashInfo{
			Path:  path2,
			Size:  0,
			Hash:  "",
			Error: nil,
		},
		Equal: false,
		Msg:   "",
	}
}

type CmpList []CmpInfo

func (r CmpList) ArrayDetail() [][]string {
	out := make([][]string, len(r)+1)
	out[0] = []string{"Path", "Size1", "Size2", "Hash1", "Hash2", "Equal", "Msg"}
	for i, v := range r {
		path := v.Hash1.Path
		if path == "" {
			path = v.Hash2.Path
		}
		out[i+1] = []string{
			path,
			fmt.Sprintf("%v", v.Hash1.Size),
			fmt.Sprintf("%v", v.Hash2.Size),
			v.Hash1.Hash,
			v.Hash2.Hash,
			fmt.Sprintf("%v", v.Equal),
			v.Msg,
		}
	}
	return out
}

func (r CmpList) Array() [][]string {
	out := make([][]string, len(r)+1)
	out[0] = []string{"EQUAL", "PATH", "MSG"}
	for i, v := range r {
		path := v.Hash1.Path
		if path == "" {
			path = v.Hash2.Path
		}
		out[i+1] = []string{
			fmt.Sprintf("%v", v.Equal),
			path,
			v.Msg,
		}
	}
	return out
}
