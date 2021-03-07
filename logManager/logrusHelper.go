package logManager

import (
	"github.com/sirupsen/logrus"
	"os"
)

// 创建一个新的logger实例。可以创建任意多个。
var log = logrus.New()

func Init() {
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")
}
