package executor

import (
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type (
	// CtrlC 捕获ctrl-c的控制器
	CtrlC struct {
		program MulServicesProgram
	}

	CtrlCOption func(*CtrlC)
)

// NewCtrlC 初始化生成CtrlC
func NewCtrlC(options ...CtrlCOption) *CtrlC {
	c := &CtrlC{}

	for _, option := range options {
		option(c)
	}

	return c
}

// WithProgram 设置程序
func WithProgram(program MulServicesProgram) CtrlCOption {
	return func(c *CtrlC) {
		c.program = program
	}
}

// 等待键盘信号
func (*CtrlC) waitSignals(signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	<-c
}

// 接收到kill信号
func (c *CtrlC) waitKill() {
	c.waitSignals(os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
}

// Run 开始运行程序，遇到os.Interrupt停止
func (c *CtrlC) Run() {
	if reflect.ValueOf(c.program).IsNil() {
		return
	}
	go func() {
		if reflect.ValueOf(c.program.Start).IsNil() {
			return
		}
		// 启动前置服务
		if err := c.program.Start(); err != nil {
			panic(err)
		}

		// 启动程序内部的服务列表
		c.startMulServices()
	}()
	c.waitKill()
	c.stopMulServices()

	if reflect.ValueOf(c.program.Stop).IsNil() {
		return
	}
	c.program.Stop()
}

// 停止应用子服务
func (c *CtrlC) startMulServices() {
	servicesSlice := c.program.ServicesRegistration()
	eg := new(errgroup.Group)
	for _, service := range servicesSlice {
		eg.Go(func() error {
			return service.Start()
		})
	}
	if err := eg.Wait(); err != nil {
		c.program.Log().Printf("service error: %s", err.Error())
	}
}

// 停止应用子服务
func (c *CtrlC) stopMulServices() {
	servicesSlice := c.program.ServicesRegistration()
	for _, service := range servicesSlice {
		service.Stop()
	}
}
