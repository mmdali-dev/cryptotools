package base64

import (
	b64 "encoding/base64"
)

func (b *BASE64) Encode(text string) string {
	var result string = b64.StdEncoding.EncodeToString([]byte(text))
	return result
}
