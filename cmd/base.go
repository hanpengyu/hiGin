package cmd

import (
	"github.com/gin-gonic/gin"
	"hi-gin/bootstrap"
)

func runServer() {
	// 初始化配置
	bootstrap.InitConfig(rootFlag)
	// 初始化日志
	bootstrap.InitLogger()

	// 启动
	engine := gin.Default()
	// 路由注册
	_ = engine.Run(":8888")
}
