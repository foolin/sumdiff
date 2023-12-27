package sumdiff_test

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/foolin/sumdiff"
	"testing"
)

func TestExample(t *testing.T) {
	path1 := "./test_data/a.txt"
	path2 := "./test_data/b.txt"

	//Compare
	ok, res, err := sumdiff.Cmp(path1, path2)
	fmt.Println("OK:", ok)
	fmt.Println("Result:", toJson(res))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Print("---------\n\n")

	//Equal
	ok, err = sumdiff.Equal(path1, path2)
	fmt.Println("OK:", ok)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Print("---------\n\n")

	//Hash
	res2, err := sumdiff.Hash(sha1.New(), path1, path2)
	fmt.Println("Result:", toJson(res2))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Print("---------\n\n")
}

func toJson(v any) string {
	data, _ := json.Marshal(v)
	return string(data)
}
