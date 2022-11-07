package executor

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	lierCmd := NewLierCmd(&LierCmdOption{
		AppName: "test",
		CmdName: "test-cmd",
		Desc:    "test-cmd-desc",
		Version: "v1.0.0",
		Author:  "biao.hu",
		Service: []Service{newtTestServer()},
	})

	ExecMulSerProgram(lierCmd)
}

func TestOption(t *testing.T) {
	lierCmd := NewLierCmd(
		NewLierCmdOption(
			WithCmdName("test-cmd-option"),
			WithDesc("test-cmd-desc"),
			WithVersion("v1.0.0"),
			WithServices(newtTestServer()),
			WithAuthor("biao.hu"),
			WithAppName("test-APP")))
	ExecMulSerProgram(lierCmd)
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
		case <-m.ch:
			println("test stop")
			return nil
		}
	}
}

func (m MyServer) Stop() {
	m.ch <- true
}

func newtTestServer() *MyServer {
	return &MyServer{}
}
