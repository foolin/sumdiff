package vo

import "fmt"

type Result struct {
	Result any `json:"result" yaml:"result"`
}

func (r Result) Array() [][]string {
	return [][]string{{"RESULT"}, {fmt.Sprintf("%v", r.Result)}}
}

func NewResult(value any) Result {
	return Result{value}
}
