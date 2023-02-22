package cipher

import "testing"

func TestAes(t *testing.T) {
	key, iv := "1234567890123456", "1234567890123456"
	aes, err := NewAes(key, iv)
	if err != nil {
		t.Error(err)
		return
	}

	encrypt, err := aes.EncryptBase64([]byte("123456"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(encrypt)

	decrypt, err := aes.DecryptBase64(encrypt)

	if err != nil {
		t.Error(err)
		return
	}
	t.Log(decrypt)

	if string(decrypt) != "123456" {
		t.Error("decrypt != 123456")
		return
	}
}
