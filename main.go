package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/logger"
	"github.com/RaymondCode/simple-demo/settings"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 抖音项目

// @title 8888组-抖音项目接口文档
// @version 1.0
// @description 字节第三届青训营抖音项目

// @contact.name 8888组全体成员
// @contact.url

// @host 127.0.0.1:8080
// @BasePath /douyin

func main() {
	//1.加载配置,从配置文件读取到viper
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err:%v\n", err)
		return
	}
	fmt.Println("Config:", settings.Conf)
	fmt.Println(settings.Conf.LogConfig == nil)

	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("Init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync() //将缓冲区日志追加到日志文件
	zap.L().Debug("init logger success...")

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) //记录日志
	initRouter(r)
	initMysql()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
