package graphql

import "testing"

func TestCheckPtr(t *testing.T) {
	if isStructPtr(nil) {
		t.Error()
	}

	if isStructPtr(struct {
	}{}) {
		t.Error()
	}

	if !isStructPtr(&struct {
	}{}) {
		t.Error()
	}
}
