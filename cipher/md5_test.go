package cipher

import "testing"

func TestMD5(t *testing.T) {
	md5Str := MD5("123456")
	if md5Str != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("md5 error")
	}
}
