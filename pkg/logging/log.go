package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filepath := getLogFileFullPath()
	F = openLogFile(filepath)
	// log.New：创建一个新的日志记录器。out定义要写入日志数据的IO句柄。prefix定义每个生成的日志行的开头。flag定义了日志记录属性

	logger = log.New(F, DefaultPrefix, log.LstdFlags)

}
func setPrefix(level Level) {
	/* 函数用于获取当前程序执行的位置（调用栈帧）的信息
		返回值包括：
	pc：uintptr类型，表示调用栈帧的程序计数器（Program Counter），用于定位函数的指令位置。
	file：字符串类型，表示源文件的完整路径。
	line：整数类型，表示代码所在的行号。
	ok：布尔值，表示是否成功获取调用栈帧信息。 */

	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		// filepath.Base()函数用于从一个路径字符串中提取文件名部分,案例:
		// path := "/path/to/file.txt"
		// filename := filepath.Base(path)
		// fmt.Println(filename) // 输出: file.txt
		logPrefix = fmt.Sprintf("[%s] [%s:%d] ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("%s", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}
