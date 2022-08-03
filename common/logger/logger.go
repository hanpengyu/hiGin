package logger

func Info(message string, fields logFields) {
	getLog().Info(message, fields)
}

//
//func Debug(message string, fields ...zap.Field) {
//	callerFields := getCallerInfoForLog()
//	fields = append(fields, callerFields...)
//	getLog().Debug(message, fields...)
//}
//
//func Error(message string, fields ...zap.Field) {
//	callerFields := getCallerInfoForLog()
//	fields = append(fields, callerFields...)
//	getLog().Error(message, fields...)
//}
//
//func Warn(message string, fields ...zap.Field) {
//	callerFields := getCallerInfoForLog()
//	fields = append(fields, callerFields...)
//	getLog().Warn(message, fields...)
//}
//
//func getCallerInfoForLog() (callerFields []zap.Field) {
//
//	pc, file, line, ok := runtime.Caller(2) // 回溯两层，拿到写日志的业务函数的信息
//	if !ok {
//		return
//	}
//	funcName := runtime.FuncForPC(pc).Name()
//	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名
//
//	callerFields = append(callerFields, zap.String("func", funcName), zap.String("file", file), zap.Int("line", line))
//	return
//}
