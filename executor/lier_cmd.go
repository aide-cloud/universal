package executor

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

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
)

// NewLierCmd 初始化生成LierCmd
func NewLierCmd(option *LierCmdOption) *LierCmd {
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
