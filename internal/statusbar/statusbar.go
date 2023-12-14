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
		MaxWidth: 150,
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

func (r *Statusbar) Show(msg string) {
	r.Message = msg
	r.once.Do(func() {
		r.doStart()
	})
}

func (r *Statusbar) doStart() {
	go func() {
		for range r.ticker.C {
			msg := r.Message
			msg = fmt.Sprintf("%v %v", time.Now().Format("2006-01-02T15:04:05"), msg)
			msg = runewidth.Truncate(msg, r.option.MaxWidth-3, "...")
			fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
			fmt.Printf("%v\r", msg)
		}
	}()
}

func (r *Statusbar) Stop() {
	r.ticker.Stop()
	fmt.Printf("%v\r", runewidth.FillLeft(" ", r.option.MaxWidth))
}
