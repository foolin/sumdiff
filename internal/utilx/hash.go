package utilx

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

func Md5(file string) (string, error) {
	return HashHex(file, md5.New())
}

func Sha1(file string) (string, error) {
	return HashHex(file, sha1.New())
}

func Sha256(file string) (string, error) {
	return HashHex(file, sha256.New())
}

func HashHex(file string, h hash.Hash) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
