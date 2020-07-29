package main

import (
	. "github.com/jaykof/chitchat/config"
	. "github.com/jaykof/chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

// 通过指定端口启动 Web 服务器
func startWebServer(port string) {
	// 在入口位置初始化全局配置
	//config := LoadConfig()
	//log.Println(config)
	r := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir(ViperConfig.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r) // 通过 router.go 中定义的路由器来分发请求

	log.Println("Starting HTTP service at " + ViperConfig.App.Address)
	err := http.ListenAndServe(ViperConfig.App.Address, nil) // 启动协程监听请求

	if err != nil {
		log.Println("An error occured starting HTTP listener at " + ViperConfig.App.Address)
		log.Println("Error: " + err.Error())
	}
}
