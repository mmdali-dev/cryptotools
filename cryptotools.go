package cryptotools

import (
	lib_base64 "github.com/mmdali-dev/cryptotools/base64"
	lib_hex "github.com/mmdali-dev/cryptotools/hex"
	lib_md5 "github.com/mmdali-dev/cryptotools/md5"
	lib_rsa "github.com/mmdali-dev/cryptotools/rsa"
)

type base64 struct{}
type hex struct{}
type md5 struct{}
type rsa struct{}

func (b *base64) Encode(text string) (result string) {
	result = lib_base64.Encode(text)
	return
}
func (b *base64) Decode(encodedText string) (result string) {
	result = lib_base64.Decode(encodedText)
	return
}

func (h *hex) Encode(text string) (result string) {
	result = lib_hex.Encode(text)
	return
}
func (h *hex) Decode(encodedText string) (result string) {
	result = lib_hex.Decode(encodedText)
	return
}

func (m *md5) Encode(text string) (result string) {
	result = lib_md5.Encode(text)
	return
}

func (r *rsa) Encode(publicKey, text string) (result string, err error) {
	result, err = lib_rsa.Encode(publicKey, text)
	return
}
func (r *rsa) Decode(privateKey, encodedText string) (result string, err error) {
	result, err = lib_rsa.Decode(privateKey, encodedText)
	return
}
func (r *rsa) GenerateKey() (publicKey string, privateKey string, err error) {
	publicKey, privateKey, err = lib_rsa.GenerateKey()
	return
}

var (
	Base64 *base64 = &base64{}
	Hex    *hex    = &hex{}
	Md5    *md5    = &md5{}
	Rsa    *rsa    = &rsa{}
)
