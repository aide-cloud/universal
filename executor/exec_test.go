package executor

import (
	"testing"
	"time"
)

func TestName(t *testing.T) {
	lierCmd := NewLierCmd(Option{
		AppName: "test",
		CmdName: "test-cmd",
		Desc:    "test-cmd-desc",
		Version: "v1.0.0",
		Author:  "biao.hu",
	})
	lierCmd.SetService(newtTestServer())
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
