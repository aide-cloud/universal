package executor

import (
	"github.com/aide-cloud/universal/alog"
	"testing"
)

func TestName(t *testing.T) {
	lierCmd := NewLierCmd()

	ExecMulSerProgram(lierCmd)
}

func TestOption(t *testing.T) {
	lierCmd := NewLierCmd(WithLogger(alog.NewLogger()), WithProperty(map[string]string{"test ": "test"}), AddProperty("test2", "test2"))
	ExecMulSerProgram(lierCmd)
}

func TestCtrlC(t *testing.T) {
	NewCtrlC(NewLierCmd(WithLogger(alog.NewLogger()))).Run()
}
