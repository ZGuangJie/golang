package main

import (
	"lastPrograma/route"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	// 路由初始化
	route.RounteInit()
	// 开启路由服务
	route.RounteStart()
	<-c
	os.Exit(0)
}
