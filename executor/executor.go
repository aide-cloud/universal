package executor

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

var globalExecutor = newCtrlC()

// ExecMulSerProgram 执行多服务程序
func ExecMulSerProgram(ex MulServicesProgram) {
	globalExecutor.servicesLock.Lock()
	globalExecutor.setMulServices(ex)
	globalExecutor.setStarter(ex)
	globalExecutor.setStopper(ex)
	globalExecutor.servicesLock.Unlock()
	globalExecutor.run()
}
