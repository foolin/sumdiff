package write

import (
	"fmt"
	"testing"
)

func TestTypeOfName(t *testing.T) {
	typ, ok := FormatOfName("json")
	fmt.Printf("%v %v", typ, ok)
}
