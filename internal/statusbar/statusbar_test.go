package statusbar

import (
	"fmt"
	"testing"
	"time"
)

func TestStatusbar_Show(t *testing.T) {
	bar := New()
	defer bar.Stop()
	millis := int((time.Second * 10).Milliseconds())
	for i := 0; i < millis; i++ {
		bar.Show(fmt.Sprintf("id=%d", i))
		time.Sleep(time.Microsecond * 200)
	}

	t.Logf("Run end!")
}
