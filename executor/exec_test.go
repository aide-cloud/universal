package executor

import (
	"github.com/aide-cloud/universal/alog"
	"testing"
	"time"
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

type MyServer struct {
	ch chan bool
}

func (m MyServer) Start() error {
	// 定时执行打印
	tick := time.NewTicker(time.Second * 1)
	count := 0
	for {
		select {
		case <-tick.C:
			println("test", count)
			count++
			//return errors.New("test error")
		case <-m.ch:
			println("test stop")
			return nil
		}
	}
}

func (m MyServer) Stop() {
	m.ch <- true
}

func (m *MyServer) Name() string {
	return "test"
}

func newtTestServer() *MyServer {
	return &MyServer{}
}
