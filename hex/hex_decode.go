package hex

import (
	HEXlib "encoding/hex"
)

func Decode(text string) string {
	result, _ := HEXlib.DecodeString(text)
	return string(result)
}
