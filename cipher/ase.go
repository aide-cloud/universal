package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type (
	AesCipher struct {
		key   []byte
		iv    []byte
		block cipher.Block
	}
)

// NewAesCipher 创建一个新的AesCipher
func NewAesCipher(key, iv string) (*AesCipher, error) {
	aesExcept := AesCipher{}
	aesExcept.key = []byte(key)
	aesExcept.iv = []byte(iv)

	var err error

	aesExcept.block, err = aes.NewCipher(aesExcept.key)
	if err != nil {
		return nil, err
	}
	return &aesExcept, nil
}

// EncryptAesBase64 加密
func (a *AesCipher) EncryptAesBase64(in string) (string, error) {
	origData := []byte(in)
	origData = pkCS5Padding(origData, a.block.BlockSize())
	crypt := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypt也可以
	bm := cipher.NewCBCEncrypter(a.block, a.iv)
	bm.CryptBlocks(crypt, origData)
	var b = base64.StdEncoding.EncodeToString(crypt)
	return b, nil
}

// DecryptAesBase64 解密
func (a *AesCipher) DecryptAesBase64(b string) (string, error) {
	crypt, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	origData := make([]byte, len(crypt))
	bm := cipher.NewCBCDecrypter(a.block, a.iv)
	bm.CryptBlocks(origData, crypt)
	origData = pkCS5UnPadding(origData)
	var out = string(origData)
	return out, nil
}

func pkCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unPadding 次
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
