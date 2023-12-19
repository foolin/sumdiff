package statusbar

import "sync"

var lock sync.Mutex
var instant *Statusbar

func Start() {
	lock.Lock()
	defer lock.Unlock()
	instant = New()
}

func Display(format string, a ...any) bool {
	lock.Lock()
	defer lock.Unlock()
	if instant == nil {
		return false
	}
	instant.Push(format, a...)
	return true
}

func Stop() bool {
	lock.Lock()
	defer lock.Unlock()
	instant.Stop()
	instant = nil
	return true
}
