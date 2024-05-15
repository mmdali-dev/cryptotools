package hex

import (
	HEXlib "encoding/hex"
)

func (h *HEX) Encode(text string) string {
	return HEXlib.EncodeToString([]byte(text))
}
