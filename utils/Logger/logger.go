package Logger

import (
	"fmt"
	"github.com/fatih/color"
	"time"
	"unsafe"
)

const (
	Danger = iota
	Warn
	Info
	Success
)

// LogInfo 日志记录
type LogInfo struct {
	Content []byte // 日志内容
	Target  string // 触发器
	Type    uint8  // 日志类型
}

// LogPos 日志产出栈
type LogPos struct {
	FilePath string // 文件路径
	FileName string // 文件名
	Line     string // 行数
}
type Logger struct {
	Data     *LogInfo
	Position *LogPos // 发生位置
	Ts       string  // 创建时间
}

// Create 消息创建
func Create() func() *Logger {
	return func() *Logger {
		logger := &Logger{
			Data:     &LogInfo{},
			Position: &LogPos{},
			Ts:       time.Now().Format("2006-01-02 15:04:05"),
		}
		buf := make([]byte, 0, 64) // test

		p := color.GreenString("[GoBot] ")
		buf = append(buf, p...)
		buf = append(buf, logger.Ts...)
		buf = append(buf, " "...)

		logger.Data.Content = buf
		return logger
	}

}
func (log *Logger) printf(p string) {
	if log.Data.Content != nil {
		log.Data.Content = append(log.Data.Content, "\t"...)
	}
	log.Data.Content = append(log.Data.Content, p...)
	log.Data.Content = append(log.Data.Content, "\n"...)
}
func (log *Logger) Warn(msg string) *Logger {
	log.printf(color.YellowString(msg))
	return log
}
func (log *Logger) Danger(msg string) *Logger {
	log.printf(color.RedString(msg))
	return log
}
func (log *Logger) Info(msg string) *Logger {
	log.printf(color.BlueString(msg))
	return log
}
func (log *Logger) Success(msg string) *Logger {
	log.printf(color.GreenString(msg))
	return log
}
func (log *Logger) Run() {
	fmt.Print(*(*string)(unsafe.Pointer(&log.Data.Content)))
}
