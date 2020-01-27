package main

import (
	"os"
	"time"
)

import (
	l4g "github.com/AlexStocks/log4go"
)

func main() {
	log := l4g.NewLogger()
	// Load the configuration (isn't this easy?)

	log.LoadConfiguration("/Users/alex/example.yml")
	defer l4g.Close() // 20190422添加此行代码，关闭默认日志输出
	defer log.Close() // 20160921添加此行代码，否则日志无法输出
	time.Sleep(2e9)   // wait send out udp package

	// And now we're ready!
	for {
		log.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
		log.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
		log.Info("About that time, eh chaps?")
		log.Warn("About that time, eh chaps?")
		log.Error("About that time, eh chaps?")
		time.Sleep(3e9)
	}

	os.Remove("trace.xml")
	os.Remove("test.log")
}
