package cmd

import (
	"fmt"
	"github.com/aide-cloud/universal/executor"
	"runtime"
	"strconv"
)

type LierCmd struct {
	AppName   string
	CmdName   string
	Version   string
	BuildTime string
	RunTime   string
	Desc      string
	Author    string
	service   []executor.Service
}

func NewCmd(appName, cmdName, desc, version, buildTime, runTime, author string) *LierCmd {
	return &LierCmd{
		AppName:   appName,
		CmdName:   cmdName,
		Desc:      desc,
		Version:   version,
		BuildTime: buildTime,
		RunTime:   runTime,
		Author:    author,
	}
}

func (cmd *LierCmd) SetService(service ...executor.Service) {
	cmd.service = service
}

// Start 启动
func (cmd *LierCmd) Start() error {
	cmd.fmtASCIIGenerator()
	return nil
}

// Stop 停止
func (cmd *LierCmd) Stop() {
	fmt.Println(fmt.Sprintf("%s-%s stoped!", cmd.AppName, cmd.CmdName))
}

// ServicesRegistration 服务注册
func (cmd *LierCmd) ServicesRegistration() []executor.Service {
	return cmd.service
}

func (cmd *LierCmd) fmtASCIIGenerator() {
	zeusStrUp := `
┌───────────────────────────────────────────────────────────────────────────────┐
│                                      _____  _____   ______                    │
│                               /\    |_   _||  __ \ |  ____|                   │
│                              /  \     | | || |  | || |__                      │
│                             / /\ \    | | || |  | ||  __|                     │
│                            / /__\ \  _| |_|| |__| || |____                    │
│                           /_/    \_\|_____||_____/ |______|                   │							
│                                 good luck and no bug                          │
└───────────────────────────────────────────────────────────────────────────────┘
`

	version := `
┌───────────────────────────────────────────────────────────────────────────────
├── app name  	: ` + cmd.AppName + `
├── cmd name  	: ` + cmd.CmdName + `
├── app desc  	: ` + cmd.Desc + `
├── app version	: ` + cmd.Version + `
├── GoVersion 	: ` + runtime.Version() + `
├── GOOS      	: ` + runtime.GOOS + `
├── NumCPU    	: ` + strconv.Itoa(runtime.NumCPU()) + `
├── RunTime    	: ` + cmd.RunTime + `
├── Date      	: ` + cmd.BuildTime + `
└───────────────────────────────────────────────────────────────────────────────
`
	fmt.Println(zeusStrUp + version)
}
