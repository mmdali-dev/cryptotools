package base64

import (
	b64 "encoding/base64"
)

func Decode(text string) string {
	result, _ := b64.StdEncoding.DecodeString(text)
	return string(result)
}
