package executor

import (
	"fmt"
	"sync"
)

type (
	LierCmd struct {
		serviceName string
		service     []Service
		property    map[string]string
		lock        sync.Mutex

		// 自定义皮肤
		skin      string
		isSetSkin bool
	}

	LierCmdOption func(*LierCmd)
)

var _ MulServicesProgram = (*LierCmd)(nil)

// NewLierCmd 初始化生成LierCmd
func NewLierCmd(options ...LierCmdOption) MulServicesProgram {
	l := &LierCmd{
		property:    make(map[string]string),
		service:     make([]Service, 0, 32),
		serviceName: "LierCmd",
	}
	l.lock.Lock()
	for _, option := range options {
		option(l)
	}
	l.lock.Unlock()
	return l
}

// Start 启动
func (cmd *LierCmd) Start() error {
	cmd.fmtASCIIGenerator()
	return nil
}

// Stop 停止
func (cmd *LierCmd) Stop() {
	fmt.Println(cmd.serviceName + " service stopped!")
}

// ServicesRegistration 服务注册
func (cmd *LierCmd) ServicesRegistration() []Service {
	return cmd.service
}

func (cmd *LierCmd) fmtASCIIGenerator() {
	fmt.Println(cmd.serviceName + " service starting...")

	if cmd.isSetSkin {
		// 自定义皮肤
		fmt.Println(cmd.skin)
		return
	}

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
┌──────────────────────────────────────────────────────────────────────────────────────`

	for k, p := range cmd.property {
		detail += fmt.Sprintf("\n├── %s: %s", k, p)
	}

	detail += `
└──────────────────────────────────────────────────────────────────────────────────────
`

	fmt.Println(detail)
}

// AddService 添加一个服务
func (cmd *LierCmd) AddService(service Service) *LierCmd {
	cmd.service = append(cmd.service, service)
	return cmd
}

// WithServices 设置服务
func WithServices(services ...Service) LierCmdOption {
	return func(l *LierCmd) {
		l.service = services
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

// WithSkin 设置皮肤
func WithSkin(skin string) LierCmdOption {
	return func(l *LierCmd) {
		l.skin = skin
		l.isSetSkin = true
	}
}

// WithServiceName 设置服务名称
func WithServiceName(serviceName string) LierCmdOption {
	return func(l *LierCmd) {
		l.serviceName = serviceName
	}
}
