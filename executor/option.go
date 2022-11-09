package executor

import (
	"github.com/aide-cloud/universal/alog"
)

type (
	// LierCmdConfig 选项
	LierCmdConfig struct {
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
		Logger  alog.Logger
	}

	LierCmdConfigFunc func(LierCmdConfig *LierCmdConfig)
)

// NewLierCmdConfig 初始化生成LierCmdConfig
func NewLierCmdConfig(opts ...LierCmdConfigFunc) *LierCmdConfig {
	option := &LierCmdConfig{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// WithServices 设置服务
func WithServices(services ...Service) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.Service = append(LierCmdConfig.Service, services...)
	}
}

// WithLogger 设置日志
func WithLogger(logger alog.Logger) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.Logger = logger
	}
}

// WithVersion 设置版本号
func WithVersion(version string) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.Version = version
	}
}

// WithAppName 设置应用名称
func WithAppName(appName string) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.AppName = appName
	}
}

// WithCmdName 设置命令名称
func WithCmdName(LierCmdName string) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.CmdName = LierCmdName
	}
}

// WithDesc 设置描述
func WithDesc(desc string) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.Desc = desc
	}
}

// WithAuthor 设置作者
func WithAuthor(author string) LierCmdConfigFunc {
	return func(LierCmdConfig *LierCmdConfig) {
		LierCmdConfig.Author = author
	}
}
