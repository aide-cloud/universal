package executor

import "github.com/aide-cloud/universal/alog"

type (
	// Service 完整的服务接口
	Service interface {
		Starter
		Stopper
		Name
	}

	// Starter 开始方法的接口
	Starter interface {
		Start() error
	}
	// Stopper 结束方法的接口
	Stopper interface {
		Stop()
	}

	// Name 服务名称
	Name interface {
		Name() string
	}

	// Logger 设置日志
	Logger interface {
		Log() alog.Logger
	}

	// MulServices 多服务程序的注册
	MulServices interface {
		ServicesRegistration() []Service
	}

	// MulServicesProgram 支持多服务启动及关闭的程序接口
	MulServicesProgram interface {
		Starter
		Stopper
		MulServices
		Logger
	}

	NewServiceFunc func(logger alog.Logger) Service
)

// ExecMulSerProgram 执行多服务程序
func ExecMulSerProgram(ex MulServicesProgram) {
	NewCtrlC(ex).Run()
}
