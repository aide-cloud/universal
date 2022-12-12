package aerror

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

const (
	ErrCodeUnknown           CodeType = 10000 + iota // 未知错误码
	ErrCodeInvalidParam                              // 参数错误
	ErrCodeInvalidToken                              // token错误
	ErrCodeInvalidSign                               // 签名错误
	ErrCodeInvalidRequest                            // 请求错误
	ErrCodeInvalidResponse                           // 响应错误
	ErrCodeInvalidData                               // 数据错误
	ErrCodeInvalidState                              // 状态错误
	ErrCodeInvalidOperation                          // 操作错误
	ErrCodeInvalidPermission                         // 权限错误
	ErrCodeInvalidUser                               // 用户错误
	ErrCodeInvalidSystem                             // 系统错误
	ErrCodeInvalidService                            // 服务错误
	ErrCodeInvalidNetwork                            // 网络错误
	ErrCodeInvalidDatabase                           // 数据库错误
	ErrCodeInvalidCache                              // 缓存错误
	ErrCodeCacheExpired                              // 缓存过期
	ErrCodeInvalidThirdParty                         // 第三方错误
)

const (
	ErrMessageUnknown           = "未知错误"
	ErrMessageInvalidParam      = "参数错误"
	ErrMessageInvalidToken      = "token错误"
	ErrMessageInvalidSign       = "签名错误"
	ErrMessageInvalidRequest    = "请求错误"
	ErrMessageInvalidResponse   = "响应错误"
	ErrMessageInvalidData       = "数据错误"
	ErrMessageInvalidState      = "状态错误"
	ErrMessageInvalidOperation  = "操作错误"
	ErrMessageInvalidPermission = "权限错误"
	ErrMessageInvalidUser       = "用户错误"
	ErrMessageInvalidSystem     = "系统错误"
	ErrMessageInvalidService    = "服务错误"
	ErrMessageInvalidNetwork    = "网络错误"
	ErrMessageInvalidDatabase   = "数据库错误"
	ErrMessageInvalidCache      = "缓存错误"
	ErrMessageCacheExpired      = "缓存过期"
	ErrMessageInvalidThirdParty = "第三方错误"
)

type (
	Coder interface {
		Code() int
	}

	Message interface {
		Message() string
	}

	HTTPStatus interface {
		HTTPStatus() int
	}

	Fields interface {
		Fields() []Field
	}

	Stacker interface {
		Stack() []byte
	}

	Error interface {
		error
		Coder
		Message
		HTTPStatus
		Fields
		Stacker
	}

	AError struct {
		httpCode int
		code     CodeType
		message  string
		fields   []Field
		stack    []byte
	}

	Option func(*AError)

	CodeType int

	Field struct {
		Key   string
		Value any
	}
)

func (a *AError) Stack() []byte {
	return a.stack
}

func (a *AError) Fields() []Field {
	return a.fields
}

func (a *AError) Message() string {
	return a.message
}

func (a *AError) HTTPStatus() int {
	return a.httpCode
}

func (a *AError) Error() string {
	return a.message
}

func (a *AError) Code() int {
	return int(a.code)
}

var _ Error = (*AError)(nil)

func New(options ...Option) Error {
	e := AError{
		httpCode: http.StatusOK,
		code:     ErrCodeUnknown,
		message:  ErrMessageUnknown,
		stack:    debug.Stack(),
	}

	for _, option := range options {
		if option != nil {
			option(&e)
		}
	}

	return &e
}

func WithCode(code CodeType) Option {
	return func(e *AError) {
		e.code = code
	}
}

func WithMessage(message string) Option {
	return func(e *AError) {
		e.message = message
	}
}

func WithErr(err error) Option {
	if err == nil {
		return nil
	}
	return func(e *AError) {
		e.message = err.Error()
	}
}

func WithStringer(str fmt.Stringer) Option {
	return func(e *AError) {
		e.message = str.String()
	}
}

func WithStack(stack []byte) Option {
	return func(e *AError) {
		e.stack = stack
	}
}

func WithHTTPStatus(httpCode int) Option {
	return func(e *AError) {
		e.httpCode = httpCode
	}
}

func WithFields(fields ...Field) Option {
	return func(e *AError) {
		e.fields = fields
	}
}

func WithCustom(code CodeType, msg string) Option {
	return func(e *AError) {
		e.code = code
		e.message = msg
	}
}
