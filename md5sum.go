package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func md5sum(name string) string {
	f, err := os.Open(name)
	if err != nil {
		return err.Error()
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%x", md5hash.Sum(nil))
}
