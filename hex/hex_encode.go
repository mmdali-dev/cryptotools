package hex

import (
	HEXlib "encoding/hex"
)

func Encode(text string) string {
	return HEXlib.EncodeToString([]byte(text))
}
