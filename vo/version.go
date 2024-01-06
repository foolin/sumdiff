package vo

type VerInfo struct {
	Version   string `json:"version" yaml:"version"`
	BuiltDate string `json:"builtDate" yaml:"builtDate"`
	Commit    string `json:"commit" yaml:"commit"`
}

func (v VerInfo) Array() [][]string {
	records := [][]string{
		{"VERSION", "BUILT_DATE", "COMMIT"},
		{v.Version, v.BuiltDate, v.Commit},
	}
	return records
}
