package sumdiff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqFile(t *testing.T) {
	for _, v := range []struct {
		expected bool
		path1    string
		path2    string
	}{
		{true, "./test_data/a.txt", "./test_data/b.txt"},
		{false, "./test_data/a.txt", "./test_data/c.txt"},
		{false, "./test_data/a.txt", "./test_data/d.txt"},
		{false, "./test_data/c.txt", "./test_data/d.txt"},
	} {
		ok, _, err := CmpFile(v.path1, v.path2)
		assert.NoError(t, err)
		assert.Equal(t, v.expected, ok, "Error at %v|%v", v.path1, v.path2)
	}
}

func TestEqDir(t *testing.T) {
	for _, v := range []struct {
		expected bool
		path1    string
		path2    string
	}{
		{true, "./test_data/data1", "./test_data/data2"},
		{false, "./test_data/data1", "./test_data/data3"},
		{false, "./test_data/data1", "./test_data/data4"},
	} {
		ok, _, err := CmpDir(v.path1, v.path2)
		assert.NoError(t, err)
		assert.Equal(t, v.expected, ok, "Error at %v|%v", v.path1, v.path2)
	}
}

func TestEq(t *testing.T) {
	for _, v := range []struct {
		expected bool
		path1    string
		path2    string
	}{
		{true, "./test_data/a.txt", "./test_data/b.txt"},
		{false, "./test_data/a.txt", "./test_data/c.txt"},
		{false, "./test_data/a.txt", "./test_data/d.txt"},
		{false, "./test_data/c.txt", "./test_data/d.txt"},
		{true, "./test_data/data1", "./test_data/data2"},
		{false, "./test_data/data1", "./test_data/data3"},
		{false, "./test_data/data1", "./test_data/data4"},
	} {
		ok, _, err := Cmp(v.path1, v.path2)
		assert.NoError(t, err)
		assert.Equal(t, v.expected, ok, "Error at %v|%v", v.path1, v.path2)
	}
}
