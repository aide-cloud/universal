package aerror

import (
	"net/http"
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

	Error interface {
		error
		Coder
		Message
		HTTPStatus
		Fields
	}

	AError struct {
		httpCode int
		code     CodeType
		message  string
		fields   []Field
	}

	Option func(*AError)

	CodeType int

	Field struct {
		Name string
		Msg  string
	}
)

var (
	ErrUnknown        = New(WithCode(ErrCodeUnknown), WithMessage(ErrMessageUnknown))
	ErrorParam        = New(WithCode(ErrCodeInvalidParam), WithMessage(ErrMessageInvalidParam))
	ErrorToken        = New(WithCode(ErrCodeInvalidToken), WithMessage(ErrMessageInvalidToken), WithHTTPStatus(http.StatusUnauthorized))
	ErrorSign         = New(WithCode(ErrCodeInvalidSign), WithMessage(ErrMessageInvalidSign), WithHTTPStatus(http.StatusUnauthorized))
	ErrorRequest      = New(WithCode(ErrCodeInvalidRequest), WithMessage(ErrMessageInvalidRequest), WithHTTPStatus(http.StatusBadRequest))
	ErrorResponse     = New(WithCode(ErrCodeInvalidResponse), WithMessage(ErrMessageInvalidResponse), WithHTTPStatus(http.StatusInternalServerError))
	ErrorData         = New(WithCode(ErrCodeInvalidData), WithMessage(ErrMessageInvalidData), WithHTTPStatus(http.StatusInternalServerError))
	ErrorState        = New(WithCode(ErrCodeInvalidState), WithMessage(ErrMessageInvalidState), WithHTTPStatus(http.StatusInternalServerError))
	ErrorOperation    = New(WithCode(ErrCodeInvalidOperation), WithMessage(ErrMessageInvalidOperation), WithHTTPStatus(http.StatusInternalServerError))
	ErrorPermission   = New(WithCode(ErrCodeInvalidPermission), WithMessage(ErrMessageInvalidPermission), WithHTTPStatus(http.StatusForbidden))
	ErrorUser         = New(WithCode(ErrCodeInvalidUser), WithMessage(ErrMessageInvalidUser), WithHTTPStatus(http.StatusUnauthorized))
	ErrorSystem       = New(WithCode(ErrCodeInvalidSystem), WithMessage(ErrMessageInvalidSystem), WithHTTPStatus(http.StatusInternalServerError))
	ErrorService      = New(WithCode(ErrCodeInvalidService), WithMessage(ErrMessageInvalidService), WithHTTPStatus(http.StatusInternalServerError))
	ErrorNetwork      = New(WithCode(ErrCodeInvalidNetwork), WithMessage(ErrMessageInvalidNetwork), WithHTTPStatus(http.StatusInternalServerError))
	ErrorDatabase     = New(WithCode(ErrCodeInvalidDatabase), WithMessage(ErrMessageInvalidDatabase), WithHTTPStatus(http.StatusInternalServerError))
	ErrorCache        = New(WithCode(ErrCodeInvalidCache), WithMessage(ErrMessageInvalidCache), WithHTTPStatus(http.StatusInternalServerError))
	ErrorCacheExpired = New(WithCode(ErrCodeCacheExpired), WithMessage(ErrMessageCacheExpired), WithHTTPStatus(http.StatusInternalServerError))
	ErrorThirdParty   = New(WithCode(ErrCodeInvalidThirdParty), WithMessage(ErrMessageInvalidThirdParty), WithHTTPStatus(http.StatusInternalServerError))
)

var (
	errorMap = map[CodeType]Error{
		ErrCodeUnknown:           ErrUnknown,
		ErrCodeInvalidParam:      ErrorParam,
		ErrCodeInvalidToken:      ErrorToken,
		ErrCodeInvalidSign:       ErrorSign,
		ErrCodeInvalidRequest:    ErrorRequest,
		ErrCodeInvalidResponse:   ErrorResponse,
		ErrCodeInvalidData:       ErrorData,
		ErrCodeInvalidState:      ErrorState,
		ErrCodeInvalidOperation:  ErrorOperation,
		ErrCodeInvalidPermission: ErrorPermission,
		ErrCodeInvalidUser:       ErrorUser,
		ErrCodeInvalidSystem:     ErrorSystem,
		ErrCodeInvalidService:    ErrorService,
		ErrCodeInvalidNetwork:    ErrorNetwork,
		ErrCodeInvalidDatabase:   ErrorDatabase,
		ErrCodeInvalidCache:      ErrorCache,
		ErrCodeCacheExpired:      ErrorCacheExpired,
		ErrCodeInvalidThirdParty: ErrorThirdParty,
	}
)

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
	}

	for _, option := range options {
		option(&e)
	}

	if _, ok := errorMap[e.code]; ok {
		return errorMap[e.code]
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
