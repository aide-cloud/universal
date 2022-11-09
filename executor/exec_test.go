package executor

import (
	"github.com/aide-cloud/universal/alog"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	lierCmd := NewLierCmd(&LierCmdConfig{
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
		NewLierCmdConfig(
			WithCmdName("test-cmd-option"),
			WithDesc("test-cmd-desc"),
			WithVersion("v1.0.0"),
			WithServices(newtTestServer()),
			WithAuthor("biao.hu"),
			WithAppName("test-APP")))
	ExecMulSerProgram(lierCmd)
}

func TestCtrlC(t *testing.T) {
	NewCtrlC(NewLierCmd(NewLierCmdConfig(WithLogger(alog.NewLogger(alog.WithOutputType(alog.OutputJsonType)))))).Run()
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
