package vo

import "fmt"

type AnyValue struct {
	Name  string `json:"name" yaml:"name"`
	Value any    `json:"value" yaml:"value"`
}

func (r AnyValue) Array() [][]string {
	if r.Name == "" {
		return [][]string{{fmt.Sprintf("%v", r.Value)}}
	}
	return [][]string{{r.Name}, {fmt.Sprintf("%v", r.Value)}}
}

func NewAnyValue(value any) AnyValue {
	return NewNameValue("", value)
}

func NewNameValue(name string, value any) AnyValue {
	return AnyValue{name, value}
}
