package executor

import (
	"github.com/aide-cloud/universal/helper/runtimehelper"
	"log"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
)

// ctrlC 捕获ctrl-c的控制器
type ctrlC struct {
	starts       Starter
	stops        Stopper
	mulServices  MulServices
	servicesLock sync.RWMutex
}

// newCtrlC 初始化生成ctrlC
func newCtrlC() *ctrlC {
	return &ctrlC{}
}

// setStarter 设置开始方法
func (c *ctrlC) setStarter(s Starter) *ctrlC {
	c.starts = s
	return c
}

// setStopper 设置结束方法
func (c *ctrlC) setStopper(s Stopper) *ctrlC {
	c.stops = s
	return c
}

// setMulServices 设置注册多服务的方法
func (c *ctrlC) setMulServices(m MulServices) *ctrlC {
	c.mulServices = m
	return c
}

// 等待键盘信号
func (*ctrlC) waitSignals(signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	<-c
}

// 接收到kill信号
func (c *ctrlC) waitKill() {
	c.waitSignals(os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
}

// run 开始运行程序，遇到os.Interrupt停止
func (c *ctrlC) run() {
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
			for index := range servicesSlice {
				go func(index int) {
					runtimehelper.Recover("service start panic")
					err := servicesSlice[index].Start()
					if err != nil {
						log.Println(err)
					}
				}(index)
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
func (c *ctrlC) stopMulServices() {
	servicesSlice := c.mulServices.ServicesRegistration()
	for _, service := range servicesSlice {
		service.Stop()
	}
}
