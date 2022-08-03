package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"runtime"
)

var _log *Logger

type logFields map[string]interface{}

type Logger struct {
	zapLogger *zap.Logger
}

func (l *Logger) Info(message string, fields logFields) {
	l.log(zap.InfoLevel, message, fields)
}

func (l *Logger) log(level zapcore.Level, message string, logFields logFields) {
	if level < zapcore.DPanicLevel && !l.zapLogger.Core().Enabled(level) {
		return
	}
	sweetenFields := l.sweetenFields(logFields)
	l.zapLogger.Info(message, sweetenFields...)
}

func (l *Logger) sweetenFields(logFields logFields) []zap.Field {
	list := make([]zap.Field, 0)
	// list = append(list, l.getCallerInfoForLog()...)
	for k, v := range logFields {
		list = append(list, zap.Any(k, v))
	}
	return list
}

func (l *Logger) getCallerInfoForLog() []zap.Field {
	var callerFields = make([]zap.Field, 0)
	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的业务函数的信息
	if !ok {
		return callerFields
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
	return callerFields
}

func (l *Logger) Sync() {
	l.zapLogger.Sync()
}

func setLogger(logger *Logger) {
	_log = logger
}

func getLog() *Logger {
	return _log
}
