package vo

type ErrInfo struct {
	Error string
}

func NewErrInfo(err error) *ErrInfo {
	return &ErrInfo{
		Error: err.Error(),
	}
}

func (r *ErrInfo) Array() [][]string {
	records := [][]string{
		{"Field", "Value"},
		{"Error", r.Error},
	}
	return records
}
