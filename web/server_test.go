package web

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer(WithServerAddr("localhost:8080"))
	fmt.Println(s)
}
