package store

import (
	"os"

	"github.com/op/go-logging"
)

// 全局日志
var Logger = logging.MustGetLogger("GoStore")

// 格式
var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
)

// level name you can refer
var LevelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// init log record
// 初始化日志
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "GoStore")
}

// 设置日志级别
// set log level
func SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "GoStore")
}

// 返回全局对象
// return global log
func Log() *logging.Logger {
	return Logger
}
