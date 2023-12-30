package write

import (
	"fmt"
	"testing"
)

func TestTypeOfName(t *testing.T) {
	typ, ok := TypeOfName("json")
	fmt.Printf("%v %v", typ, ok)
}
