package web

import "testing"

func TestLierWebGin(t *testing.T) {
	myWebServer := NewGin()
	err := myWebServer.Start()
	if err != nil {
		t.Error(err)
	}
}
