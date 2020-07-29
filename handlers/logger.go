package handlers

import (
	. "github.com/jaykof/chitchat/config"
	"log"
	"os"
)

var logger *log.Logger

// 每个包都可以一个或多个init()函数，它们会按照包的导入顺序和文件编译顺序被顺序执行
func init() {
	file, err := os.OpenFile(ViperConfig.App.Log+"/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 为什么不命名为 error?避免和 error 类型重名
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}
