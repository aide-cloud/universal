package executor

type (
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

	OptionFunc func(option *Option)
)

func NewOption(opts ...OptionFunc) *Option {
	option := &Option{}
	for _, opt := range opts {
		opt(option)
	}
	return option
}

// WithVersion 设置版本号
func WithVersion(version string) OptionFunc {
	return func(option *Option) {
		option.Version = version
	}
}

// WithAppName 设置应用名称
func WithAppName(appName string) OptionFunc {
	return func(option *Option) {
		option.AppName = appName
	}
}

// WithCmdName 设置命令名称
func WithCmdName(cmdName string) OptionFunc {
	return func(option *Option) {
		option.CmdName = cmdName
	}
}

// WithDesc 设置描述
func WithDesc(desc string) OptionFunc {
	return func(option *Option) {
		option.Desc = desc
	}
}

// WithAuthor 设置作者
func WithAuthor(author string) OptionFunc {
	return func(option *Option) {
		option.Author = author
	}
}
