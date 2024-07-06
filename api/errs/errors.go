package errs

import (
	"github.com/jongyunha/advance-go-web-application/api/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ApplicationError struct {
	statusCode int
	code       ErrorCode
	message    string
	error      error
	level      zapcore.Level
}

func newError(statusCode int, code ErrorCode, message string, level zapcore.Level) *ApplicationError {
	return &ApplicationError{
		statusCode: statusCode,
		code:       code,
		message:    message,
		level:      level,
	}
}

func NewWarnError(statusCode int, code ErrorCode, message string) *ApplicationError {
	return newError(statusCode, code, message, zapcore.WarnLevel)
}

func NewError(statusCode int, code ErrorCode, message string) *ApplicationError {
	return newError(statusCode, code, message, zapcore.ErrorLevel)
}

func (e *ApplicationError) StatusCode() int {
	return e.statusCode
}

func (e *ApplicationError) Error() string {
	return e.message
}

func (e *ApplicationError) Level() zapcore.Level {
	return e.level
}

func (e *ApplicationError) Code() ErrorCode {
	return e.code
}

func (e *ApplicationError) Message() string {
	return e.message
}

func (e *ApplicationError) Log() {
	core.Logger.Error(e.message, zap.Error(e.error))
}
