package executor

import (
	"fmt"
	"github.com/aide-cloud/universal/helper/runtimehelper"
	"log"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type (
	// Service 完整的服务接口
	Service interface {
		Starter
		Stopper
	}

	// Starter 开始方法的接口
	Starter interface {
		Start() error
	}
	// Stopper 结束方法的接口
	Stopper interface {
		Stop()
	}

	// MulServices 多服务程序的注册
	MulServices interface {
		ServicesRegistration() []Service
	}

	// Program 完整程序的接口
	Program interface {
		Starter
		Stopper
	}

	// MulServicesProgram 支持多服务启动及关闭的程序接口
	MulServicesProgram interface {
		Starter
		Stopper
		MulServices
	}
)

// ctrlC 捕获ctrl-c的控制器
type ctrlC struct {
	starts       Starter
	stops        Stopper
	mulServices  MulServices
	servicesLock sync.RWMutex
}

// NewCtrlC 初始化生成ctrlC
func NewCtrlC() *ctrlC {
	return &ctrlC{}
}

// SetStarter 设置开始方法
func (c *ctrlC) SetStarter(s Starter) *ctrlC {
	c.starts = s
	return c
}

// SetStopper 设置结束方法
func (c *ctrlC) SetStopper(s Stopper) *ctrlC {
	c.stops = s
	return c
}

// SetMulServices 设置注册多服务的方法
func (c *ctrlC) SetMulServices(m MulServices) *ctrlC {
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

// Run 开始运行程序，遇到os.Interrupt停止
func (c *ctrlC) Run() {
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

type (
	LierCmd struct {
		appName string
		cmdName string
		version string
		runTime string
		desc    string
		author  string
		service []Service
	}

	// Option 选项
	Option struct {
		// AppName 应用名称
		AppName string
		// CmdName 命令名称
		CmdName string
		// Version 版本号
		Version string
		// Desc 描述
		Desc string
		// Author 作者
		Author string
	}
)

// NewLierCmd 初始化生成LierCmd
func NewLierCmd(option Option) *LierCmd {
	return &LierCmd{
		appName: option.AppName,
		cmdName: option.CmdName,
		desc:    option.Desc,
		version: option.Version,
		runTime: time.Now().Format("2006-01-02 15:04:05"),
		author:  option.Author,
	}
}

func (cmd *LierCmd) SetService(service ...Service) {
	cmd.service = service
}

// Start 启动
func (cmd *LierCmd) Start() error {
	cmd.fmtASCIIGenerator()
	return nil
}

// Stop 停止
func (cmd *LierCmd) Stop() {
	fmt.Println(fmt.Sprintf("%s-%s stoped!", cmd.appName, cmd.cmdName))
}

// ServicesRegistration 服务注册
func (cmd *LierCmd) ServicesRegistration() []Service {
	return cmd.service
}

func (cmd *LierCmd) fmtASCIIGenerator() {
	zeusStrUp := `
┌───────────────────────────────────────────────────────────────────────────────────────┐
│                                      _____  _____   ______                            │
│                               /\    |_   _||  __ \ |  ____|                           │
│                              /  \     | | || |  | || |__                              │
│                             / /\ \    | | || |  | ||  __|                             │
│                            / /__\ \  _| |_|| |__| || |____                            │
│                           /_/    \_\|_____||_____/ |______|                           │							
│                                 good luck and no bug                                  │
└───────────────────────────────────────────────────────────────────────────────────────┘
`

	version := `
┌───────────────────────────────────────────────────────────────────────────────────────
├── app name  	: ` + cmd.appName + `
├── cmd name  	: ` + cmd.cmdName + `
├── app desc  	: ` + cmd.desc + `
├── app version	: ` + cmd.version + `
├── GoVersion 	: ` + runtime.Version() + `
├── GOOS      	: ` + runtime.GOOS + `
├── NumCPU    	: ` + strconv.Itoa(runtime.NumCPU()) + `
├── RunTime    	: ` + cmd.runTime + `
└───────────────────────────────────────────────────────────────────────────────────────
`
	fmt.Println(zeusStrUp + version)
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
