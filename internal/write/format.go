package write

import "strings"

type Format int

const (
	None Format = iota
	Table
	Json
	Csv
	Yaml
)

var formatNames = []string{"none", "table", "json", "csv", "yaml"}

// String - Name
func (t Format) String() string {
	return formatNames[t]
}

// Index - Index value
func (t Format) Index() int {
	return int(t)
}

func FormatOfName(name string) (Format, bool) {
	if name == "" {
		return None, false
	}
	fmtName := strings.ToLower(name)
	for i, v := range formatNames {
		if fmtName == v {
			return FormatOfValue(i)
		}
	}
	return None, false
}

func FormatOfValue(value int) (Format, bool) {
	return Format(value), true
}
