package md5

import (
	Md5 "crypto/md5"
	"fmt"
)

type MD5 struct{}

func (m *MD5) Encode(text string) string {
	var result [16]byte = Md5.Sum([]byte(text))
	return fmt.Sprintf("%x", result)
}
