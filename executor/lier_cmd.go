package executor

import (
	"fmt"
	"github.com/aide-cloud/universal/alog"
	"sync"
)

type (
	LierCmd struct {
		service  []Service
		logger   alog.Logger
		property map[string]string
		lock     sync.Mutex
	}

	LierCmdOption func(*LierCmd)
)

var _ MulServicesProgram = (*LierCmd)(nil)

// NewLierCmd 初始化生成LierCmd
func NewLierCmd(options ...LierCmdOption) MulServicesProgram {
	l := &LierCmd{
		property: make(map[string]string),
		service:  make([]Service, 0, 32),
	}
	l.lock.Lock()
	for _, option := range options {
		option(l)
	}
	l.lock.Unlock()
	return l
}

func (cmd *LierCmd) Log() alog.Logger {
	if cmd.logger == nil {
		cmd.logger = alog.NewLogger()
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
	cmd.Log().Warn("master service stopped!")
}

// ServicesRegistration 服务注册
func (cmd *LierCmd) ServicesRegistration() []Service {
	return cmd.service
}

func (cmd *LierCmd) fmtASCIIGenerator() {
	fmt.Println(`┌───────────────────────────────────────────────────────────────────────────────────────┐
│                                      _____  _____   ______                            │
│                               /\    |_   _||  __ \ |  ____|                           │
│                              /  \     | | || |  | || |__                              │
│                             / /\ \    | | || |  | ||  __|                             │
│                            / /__\ \  _| |_|| |__| || |____                            │
│                           /_/    \_\|_____||_____/ |______|                           │							
│                                 good luck and no bug                                  │
└───────────────────────────────────────────────────────────────────────────────────────┘`)

	if cmd.property == nil || len(cmd.property) == 0 {
		return
	}

	detail := `
┌───────────────────────────────────────────────────────────────────────────────────────`

	for k, p := range cmd.property {
		detail += fmt.Sprintf("\n├── %s: %s", k, p)
	}

	detail += `
└───────────────────────────────────────────────────────────────────────────────────────
`

	fmt.Println(detail)
}

// WithServices 设置服务
func WithServices(services ...NewServiceFunc) LierCmdOption {
	return func(l *LierCmd) {
		for _, service := range services {
			l.service = append(l.service, service(l.logger))
		}
	}
}

// AddService 添加一个服务
func AddService(service NewServiceFunc) LierCmdOption {
	return func(l *LierCmd) {
		if service == nil {
			return
		}
		l.service = append(l.service, service(l.logger))
	}
}

// WithLogger 设置日志
func WithLogger(logger alog.Logger) LierCmdOption {
	return func(l *LierCmd) {
		l.logger = logger
	}
}

// WithProperty 设置属性
func WithProperty(property map[string]string) LierCmdOption {
	return func(l *LierCmd) {
		if property == nil || len(property) == 0 {
			return
		}
		l.property = property
	}
}

// AddProperty 添加属性
func AddProperty(key, value string) LierCmdOption {
	return func(l *LierCmd) {
		if key == "" {
			return
		}
		l.property[key] = value
	}
}
