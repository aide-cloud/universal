package cipher

import "testing"

func TestAes(t *testing.T) {
	key, iv := "1234567890123456", "1234567890123456"
	aes, err := NewAesCipher(key, iv)
	if err != nil {
		t.Error(err)
		return
	}

	encrypt, err := aes.EncryptAesBase64("123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(encrypt)

	decrypt, err := aes.DecryptAesBase64(encrypt)

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(decrypt)

	if decrypt != "123456" {
		t.Error("decrypt != 123456")
		return
	}
}
