package executor

import (
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type (
	// CtrlC 捕获ctrl-c的控制器
	CtrlC struct {
		program MulServicesProgram
	}
)

// NewCtrlC 初始化生成CtrlC
func NewCtrlC(ex MulServicesProgram) *CtrlC {
	return &CtrlC{
		program: ex,
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
	go func() {
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