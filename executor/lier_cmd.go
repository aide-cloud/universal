package executor

import (
	"log"
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
		logger  *log.Logger
	}
)

var _ MulServicesProgram = (*LierCmd)(nil)

// NewLierCmd 初始化生成LierCmd
func NewLierCmd(cfg *LierCmdConfig) *LierCmd {
	return &LierCmd{
		appName: cfg.AppName,
		cmdName: cfg.CmdName,
		desc:    cfg.Desc,
		version: cfg.Version,
		runTime: time.Now().Format("2006-01-02 15:04:05"),
		author:  cfg.Author,
		service: cfg.Service,
		logger:  cfg.Logger,
	}
}

func (cmd *LierCmd) Log() *log.Logger {
	if cmd.logger == nil {
		cmd.logger = log.New(log.Writer(), cmd.appName+"-"+cmd.cmdName+" ", log.Flags())
	}
	return cmd.logger
}

// Start 启动
func (cmd *LierCmd) Start() error {
	cmd.fmtASCIIGenerator()
	return nil
}

// Stop 停止
func (cmd *LierCmd) Stop() {
	cmd.Log().Printf("%s-%s stoped!\n", cmd.appName, cmd.cmdName)
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
	cmd.Log().Println(zeusStrUp + version)
}
