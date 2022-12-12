package aerror

import (
	"testing"
)

func TestNew(t *testing.T) {
	err := New(WithMessage("test"))
	if err == nil {
		t.Fatal("err is nil")
	}

	if err.Error() != "test" {
		t.Fatal("err is not test")
	}

	if err.HTTPStatus() != 200 {
		t.Fatal("err http status is not 200")
	}

	t.Log(string(err.Stack()))
}
