package hex

import (
	HEXlib "encoding/hex"
)

func (h *HEX) Decode(text string) string {
	result, _ := HEXlib.DecodeString(text)
	return string(result)
}
