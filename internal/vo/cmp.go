package vo

import "fmt"

type CmpVo struct {
	X   *HashVo `json:"x"`
	Y   *HashVo `json:"y"`
	OK  bool    `json:"ok"`
	Msg string  `json:"msg"`
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
		OK:  false,
		Msg: "",
	}
}

func CmpToTable(list []*CmpVo) [][]string {
	out := make([][]string, len(list)+1)
	out[0] = []string{"Path", "Size1", "Size2", "Hash1", "Hash2", "OK", "Msg"}
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
			fmt.Sprintf("%v", v.OK),
			v.Msg,
		}
	}
	return out
}
