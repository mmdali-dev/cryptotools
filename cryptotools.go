package cryptotools

import (
	"github.com/mmdali-dev/cryptotools/base64"
	"github.com/mmdali-dev/cryptotools/hex"
	"github.com/mmdali-dev/cryptotools/md5"
	"github.com/mmdali-dev/cryptotools/rsa"
)

var (
	MD5    *md5.MD5       = &md5.MD5{}
	RSA    *rsa.RSA       = &rsa.RSA{}
	HEX    *hex.HEX       = &hex.HEX{}
	BASE64 *base64.BASE64 = &base64.BASE64{}
)
