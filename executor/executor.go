package executor

import (
	"log"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
)

// CtrlC 捕获ctrl-c的控制器
type CtrlC struct {
	starts       Starter
	stops        Stopper
	mulServices  MulServices
	servicesLock sync.RWMutex
}

// NewCtrlC 初始化生成CtrlC
func NewCtrlC() *CtrlC {
	return &CtrlC{}
}

// SetStarter 设置开始方法
func (c *CtrlC) SetStarter(s Starter) *CtrlC {
	c.starts = s
	return c
}

// SetStopper 设置结束方法
func (c *CtrlC) SetStopper(s Stopper) *CtrlC {
	c.stops = s
	return c
}

// SetMulServices 设置注册多服务的方法
func (c *CtrlC) SetMulServices(m MulServices) *CtrlC {
	c.mulServices = m
	return c
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
		if reflect.ValueOf(c.starts).IsNil() {
			return
		}
		// 启动前置服务
		if err := c.starts.Start(); err != nil {
			panic(err)
		}

		// 启动程序内部的服务列表
		if c.mulServices != nil {
			servicesSlice := c.mulServices.ServicesRegistration()
			for _, service := range servicesSlice {
				err := service.Start()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
	c.waitKill()
	c.stopMulServices()

	if reflect.ValueOf(c.stops).IsNil() {
		return
	}
	c.stops.Stop()
}

// 停止应用子服务
func (c *CtrlC) stopMulServices() {
	servicesSlice := c.mulServices.ServicesRegistration()
	for _, service := range servicesSlice {
		service.Stop()
	}
}

var globalExecutor = NewCtrlC()

// ExecMulSerProgram 执行多服务程序
func ExecMulSerProgram(ex MulServicesProgram) {
	globalExecutor.servicesLock.Lock()
	globalExecutor.SetMulServices(ex)
	globalExecutor.SetStarter(ex)
	globalExecutor.SetStopper(ex)
	globalExecutor.servicesLock.Unlock()
	globalExecutor.Run()
}
