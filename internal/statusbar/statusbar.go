package statusbar

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"sync"
	"time"
)

type Statusbar struct {
	Message string
	once    sync.Once
	ticker  *time.Ticker
	option  *Option
}

type Option struct {
	Interval time.Duration
	MaxWidth int
}

func DefaultOption() *Option {
	return &Option{
		Interval: time.Second,
		MaxWidth: 100,
	}
}

func New(optFns ...func(option *Option)) *Statusbar {
	opt := DefaultOption()
	for _, fn := range optFns {
		fn(opt)
	}

	return &Statusbar{
		Message: "",
		once:    sync.Once{},
		ticker:  time.NewTicker(opt.Interval),
		option:  opt,
	}
}

func (r *Statusbar) Display(format string, a ...any) {
	r.Message = fmt.Sprintf(format, a...)
	r.once.Do(func() {
		r.doStart()
	})
}

func (r *Statusbar) doStart() {
	go func() {
		for range r.ticker.C {
			msg := r.Message
			msg = fmt.Sprintf("%v %v", time.Now().Format("15:04:05"), msg)
			msg = truncateMid(msg, r.option.MaxWidth, "...")
			if len(msg) > 0 {
				fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
			}
			fmt.Printf("%v\r", msg)
		}
	}()
}

func (r *Statusbar) Stop() {
	r.ticker.Stop()
	fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
}

func Display(r *Statusbar, format string, a ...any) bool {
	if r == nil {
		return false
	}
	r.Display(format, a...)
	return true
}

var strwidth = &runewidth.Condition{
	EastAsianWidth:     true,
	StrictEmojiNeutral: true,
}

func truncateMid(s string, width int, ellipsis string) string {
	keepLeft := width / 2
	size := strwidth.StringWidth(s)
	if size <= width {
		return s
	}
	if keepLeft <= 0 || keepLeft >= size {
		return strwidth.Truncate(s, width, ellipsis)
	}
	lstr := strwidth.Truncate(s, keepLeft, ellipsis)
	rstr := strwidth.TruncateLeft(s, size-(width-keepLeft), "")
	return lstr + rstr
}
