package main

import (
	"fmt"
	"github.com/aide-cloud/universal/cipher"
)

func show(str string) {
	md5Str := cipher.MD5(str)
	fmt.Println(md5Str)
}

func main() {
	show("123")
	show("abc")
	show("xxx")
}
