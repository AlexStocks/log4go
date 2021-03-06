package main

import (
	"time"
)

import (
	l4g "github.com/AlexStocks/log4go"
)

func main() {
	// 下面 两行 等效
	// log.NewDefaultLogger(log.INFO).SetAsDefaultLogger()
	// log.SetLogLevel(log.INFO)

	log := l4g.NewLogger()
	defer log.Close()
	defer l4g.Close() // 关闭默认的全局

	log.AddFilter("stdout", l4g.DEBUG, l4g.NewConsoleLogWriter(true))
	log.Info("The time is now: %s", time.Now().Format("15:04:05 MST 2006/01/02"))
}
