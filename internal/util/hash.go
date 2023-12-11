package util

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

func Md5(file string) (string, error) {
	return HashHex(md5.New(), file)
}

func Sha1(file string) (string, error) {
	return HashHex(sha1.New(), file)
}

func Sha256(file string) (string, error) {
	return HashHex(sha256.New(), file)
}

func Sha512(file string) (string, error) {
	return HashHex(sha512.New(), file)
}

func HashHex(h hash.Hash, file string) (string, error) {
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
