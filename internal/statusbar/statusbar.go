package statusbar

import (
	"fmt"
	"sync"
	"time"

	"github.com/mattn/go-runewidth"
)

type State int

const (
	StatusNone    State = 0
	StatusRunning State = 1
	StatusStop    State = -1
)

type Statusbar struct {
	Message string
	once    sync.Once
	ticker  *time.Ticker
	status  State //0-none, 1=running -1=stop
	option  *Option
}

type Option struct {
	Interval time.Duration
	MaxWidth int
	Disable  bool
}

func DefaultOption() *Option {
	return &Option{
		Interval: time.Millisecond * 200,
		MaxWidth: 100,
		Disable:  false,
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
		status:  StatusNone,
		option:  opt,
	}
}

func (r *Statusbar) Push(format string, a ...any) {
	r.Message = fmt.Sprintf(format, a...)
	r.once.Do(func() {
		r.doStart()
	})
}

func (r *Statusbar) Display(format string, a ...any) {
	r.Message = fmt.Sprintf(format, a...)
	r.display(r.Message)
}

func (r *Statusbar) DisplayEnd(format string, a ...any) {
	r.Display(format, a...)
	fmt.Println()
}

func (r *Statusbar) Stop() {
	r.status = StatusStop
	r.ticker.Stop()
	r.Clean()
}

func (r *Statusbar) Clean() {
	if r.option.Disable {
		return
	}
	fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
}

func (r *Statusbar) doStart() {
	r.status = StatusRunning
	go func() {
		for {
			if r.status == -1 {
				break
			}
			select {
			case <-r.ticker.C:
				r.display(r.Message)
			case <-time.After(time.Millisecond):
				r.display(r.Message)
			}
		}
	}()
}

func (r *Statusbar) display(msg string) {
	if r.option.Disable {
		return
	}
	msg = fmt.Sprintf("%v %v", time.Now().Format("15:04:05"), msg)
	msg = truncateMid(msg, r.option.MaxWidth, "...")
	if len(msg) > 0 {
		fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
	}
	fmt.Printf("%v\r", msg)
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
