package executor

import (
	"testing"
)

func TestName(t *testing.T) {
	lierCmd := NewLierCmd()

	ExecMulSerProgram(lierCmd)
}

func TestOption(t *testing.T) {
	lierCmd := NewLierCmd(WithProperty(map[string]string{"test ": "test"}))
	ExecMulSerProgram(lierCmd)
}

func TestCtrlC(t *testing.T) {
	NewCtrlC(NewLierCmd()).Run()
}
