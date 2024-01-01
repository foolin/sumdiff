package vo

import "fmt"

type Any struct {
	Name  string
	Value any
}

func (r Any) Array() [][]string {
	return [][]string{{r.Name}, {fmt.Sprintf("%v", r.Value)}}
}

func NewAny(name string, value any) Any {
	return Any{name, value}
}

func NewResult(value any) Any {
	return NewAny("RESULT", value)
}
