package sumdiff

import (
	"github.com/foolin/sumdiff/internal/statusbar"
	"time"
)

type Option struct {
	StatusbarWidth    int
	StatusbarInterval time.Duration
	StatusbarDisable  bool
}

func DefaultOption() *Option {
	sbarOpt := statusbar.DefaultOption()
	return &Option{
		StatusbarWidth:    sbarOpt.MaxWidth,
		StatusbarDisable:  sbarOpt.Disable,
		StatusbarInterval: sbarOpt.Interval,
	}
}
