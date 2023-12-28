package write

import "strings"

type Type int

const (
	None Type = iota
	Table
	Json
	Csv
	Yaml
)

var typeNames = []string{"none", "table", "json", "csv", "yaml"}

// String - Name
func (t Type) String() string {
	return typeNames[t]
}

// Index - Index value
func (t Type) Index() int {
	return int(t)
}

func TypeOfName(name string) (Type, bool) {
	if name == "" {
		return None, false
	}
	fmtName := strings.ToLower(name)
	for i, v := range typeNames {
		if fmtName == v {
			return TypeOfValue(i)
		}
	}
	return None, false
}

func TypeOfValue(value int) (Type, bool) {
	return Type(value), true
}
