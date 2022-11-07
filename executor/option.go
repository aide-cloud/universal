package executor

import "log"

type (
	// LierCmdOption 选项
	LierCmdOption struct {
		// AppName 应用名称
		AppName string
		// LierCmdName 命令名称
		CmdName string
		// Version 版本号
		Version string
		// Desc 描述
		Desc string
		// Author 作者
		Author string

		Service []Service
		Logger  *log.Logger
	}

	LierCmdOptionFunc func(LierCmdOption *LierCmdOption)
)

// NewLierCmdOption 初始化生成LierCmdOption
func NewLierCmdOption(opts ...LierCmdOptionFunc) *LierCmdOption {
	option := &LierCmdOption{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// WithServices 设置服务
func WithServices(services ...Service) LierCmdOptionFunc {
	return func(lierCmdOption *LierCmdOption) {
		lierCmdOption.Service = append(lierCmdOption.Service, services...)
	}
}

// WithLogger 设置日志
func WithLogger(logger *log.Logger) LierCmdOptionFunc {
	return func(lierCmdOption *LierCmdOption) {
		lierCmdOption.Logger = logger
	}
}

// WithVersion 设置版本号
func WithVersion(version string) LierCmdOptionFunc {
	return func(LierCmdOption *LierCmdOption) {
		LierCmdOption.Version = version
	}
}

// WithAppName 设置应用名称
func WithAppName(appName string) LierCmdOptionFunc {
	return func(LierCmdOption *LierCmdOption) {
		LierCmdOption.AppName = appName
	}
}

// WithCmdName 设置命令名称
func WithCmdName(LierCmdName string) LierCmdOptionFunc {
	return func(LierCmdOption *LierCmdOption) {
		LierCmdOption.CmdName = LierCmdName
	}
}

// WithDesc 设置描述
func WithDesc(desc string) LierCmdOptionFunc {
	return func(LierCmdOption *LierCmdOption) {
		LierCmdOption.Desc = desc
	}
}

// WithAuthor 设置作者
func WithAuthor(author string) LierCmdOptionFunc {
	return func(LierCmdOption *LierCmdOption) {
		LierCmdOption.Author = author
	}
}
