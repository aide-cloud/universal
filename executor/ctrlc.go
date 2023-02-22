package executor

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type (
	// CtrlC 捕获ctrl-c的控制器
	CtrlC struct {
		program MulServicesProgram
		// 信号通道
		signalChan chan os.Signal
	}
)

// NewCtrlC 初始化生成CtrlC
func NewCtrlC(ex MulServicesProgram) *CtrlC {
	return &CtrlC{
		program:    ex,
		signalChan: make(chan os.Signal, 1),
	}
}

// 等待键盘信号
func (c *CtrlC) waitSignals(signals ...os.Signal) {
	signal.Notify(c.signalChan, signals...)
	<-c.signalChan
}

// 接收到kill信号
func (c *CtrlC) waitKill() {
	c.waitSignals(os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
}

// Run 开始运行程序，遇到os.Interrupt停止
func (c *CtrlC) Run() {
	go func() {
		defer c.recover()
		// 启动前置服务
		if err := c.program.Start(); err != nil {
			panic(err)
		}

		// 启动程序内部的服务列表
		c.startMulServices()
	}()

	c.waitKill()
	c.stopMulServices()
	c.program.Stop()
}

func (c *CtrlC) recover() {
	if err := recover(); err != nil {
		fmt.Println(err)
		c.signalChan <- os.Kill
	}
}

// 停止应用子服务
func (c *CtrlC) startMulServices() {
	servicesSlice := c.program.ServicesRegistration()
	for _, service := range servicesSlice {
		go func(s Service) {
			defer c.recover()
			if err := s.Start(); err != nil {
				panic(err)
			}
		}(service)
	}
}

// 停止应用子服务
func (c *CtrlC) stopMulServices() {
	servicesSlice := c.program.ServicesRegistration()
	for _, service := range servicesSlice {
		service.Stop()
	}
}
