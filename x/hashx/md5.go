package hashx

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5(r io.Reader) (string, error) {
	h := md5.New()

	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
