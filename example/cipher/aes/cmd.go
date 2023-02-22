package main

import (
	"fmt"
	"github.com/aide-cloud/universal/cipher"
)

func main() {
	key, iv := "1234567890123456", "1234567890123456"
	aes, err := cipher.NewAes(key, iv)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 加密
	encryptStr, err := aes.EncryptAesBase64("123")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解密
	decryptStr, err := aes.DecryptAesBase64(encryptStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("加密前：", "123")
	fmt.Println("加密后：", encryptStr)
	fmt.Println("解密后：", decryptStr)
}
