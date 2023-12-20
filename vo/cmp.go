package vo

import "fmt"

type CmpVo struct {
	Hash1 *HashVo `json:"hash1"`
	Hash2 *HashVo `json:"hash2"`
	Equal bool    `json:"equal"`
	Msg   string  `json:"msg"`
}

func NewCmpVo(path1, path2 string) *CmpVo {
	return &CmpVo{
		Hash1: &HashVo{
			Path:  path1,
			Size:  0,
			Hash:  "",
			Error: nil,
		},
		Hash2: &HashVo{
			Path:  path2,
			Size:  0,
			Hash:  "",
			Error: nil,
		},
		Equal: false,
		Msg:   "",
	}
}

func CmpToTable(list []*CmpVo) [][]string {
	out := make([][]string, len(list)+1)
	out[0] = []string{"Path", "Size1", "Size2", "Hash1", "Hash2", "Equal", "Msg"}
	for i, v := range list {
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

func CmpToLiteTable(list []*CmpVo) [][]string {
	out := make([][]string, len(list)+1)
	out[0] = []string{"Path", "Equal", "Msg"}
	for i, v := range list {
		path := v.Hash1.Path
		if path == "" {
			path = v.Hash2.Path
		}
		out[i+1] = []string{
			path,
			fmt.Sprintf("%v", v.Equal),
			v.Msg,
		}
	}
	return out
}
