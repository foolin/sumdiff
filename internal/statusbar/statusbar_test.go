package statusbar

import (
	"crypto/sha512"
	"fmt"
	"testing"
	"time"
)

func TestStatusbar_Show(t *testing.T) {
	bar := New()
	defer bar.Stop()
	millis := int((time.Second * 10).Milliseconds())
	hash := sha512.New()
	for i := 0; i < millis; i++ {
		hash.Reset()
		hash.Write([]byte(fmt.Sprintf("%v", i)))
		bar.Display(fmt.Sprintf("id=%d %x", i, hash.Sum(nil)))
		time.Sleep(time.Microsecond * 200)
	}

	t.Logf("Run end!")
}

func TestStatusbar_Show2(t *testing.T) {
	bar := New()
	defer bar.Stop()
	seconds := 20
	hash := sha512.New()
	for i := 0; i < seconds; i++ {
		hash.Reset()
		hash.Write([]byte(fmt.Sprintf("%v", i)))
		bar.Display(fmt.Sprintf("id=%d %x", i, hash.Sum(nil)))
		time.Sleep(time.Second * 2)
	}

	t.Logf("Run end!")
}

func TestTruncateMid(t *testing.T) {
	str := "States that a function type does not have ellipsis parameter."
	str2 := "[Github] 提醒 -k8s 合并请求：状态变更，状态栏导致显示错误的问题！"
	t.Log(truncateMid(str, 50, "..."))
	t.Log(truncateMid(str2, 50, "..."))
}
