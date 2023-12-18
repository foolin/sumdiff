package vo

import "fmt"

type CmpVo struct {
	X     *HashVo `json:"x"`
	Y     *HashVo `json:"y"`
	Equal bool    `json:"equal"`
	Msg   string  `json:"msg"`
}

func NewCmpVo(path1, path2 string) *CmpVo {
	return &CmpVo{
		X: &HashVo{
			Path:  path1,
			Size:  0,
			Hash:  "",
			Error: nil,
		},
		Y: &HashVo{
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
		path := v.X.Path
		if path == "" {
			path = v.Y.Path
		}
		out[i+1] = []string{
			path,
			fmt.Sprintf("%v", v.X.Size),
			fmt.Sprintf("%v", v.Y.Size),
			v.X.Hash,
			v.Y.Hash,
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
		path := v.X.Path
		if path == "" {
			path = v.Y.Path
		}
		out[i+1] = []string{
			path,
			fmt.Sprintf("%v", v.Equal),
			v.Msg,
		}
	}
	return out
}
