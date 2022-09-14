package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	errorPrefix = "[ERROR]"
	infoPrefix = "[INFO]"
	warnPrefix = "[WARN]"
	debugPrefix = "[DEBUG]"
)

func Error(format string, objects ...interface{})  {
	printLog(errorPrefix, format, objects...)
}
func Info(format string, objects ...interface{})  {
	printLog(infoPrefix, format, objects...)
}
func Warn(format string, objects ...interface{})  {
	printLog(warnPrefix, format, objects...)
}
func Debug(format string, objects ...interface{})  {
	printLog(debugPrefix, format, objects...)
}

func printLog(logType, format string, objects ...interface{})  {
	funcName, file, line, ok := runtime.Caller(2) //弹出两层函数栈再打印
	var prefix string
	if ok {
		file = file[strings.LastIndex(file, "/")+1:]
		prefix = fmt.Sprintf("%-9s [%-19s]   File:[%s:%d]   Func:[%s] --- ", logType, time.Now().Format("2006-01-02 15:04:05"), file, line, runtime.FuncForPC(funcName).Name())
	}else{
		prefix = fmt.Sprintf("%-9s [%-19s]   ", logType, time.Now().Format("2006-01-02 15:04:05"))
	}
	// fmt.Printf( prefix + format + "\n", objects...)
	fmt.Fprintf(os.Stdout, prefix + format + "\n", objects...)
}
