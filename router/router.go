package router

import "github.com/gin-gonic/gin"

func SetRouter() *gin.Engine {
	r := gin.Default()
	// 将asset/目录映射到/asset接口路径，提供静态文件访问功能
	r.Static("/asset", "asset/")
	// 将/favicon.ico路由直接映射到asset/images/favicon.ico这个具体的文件，用于提供单个静态文件的访问
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	// 加载views目录下所有子目录的HTML模板，用于Gin的HTML渲染
	r.LoadHTMLGlob("views/**/*")
	return r
}
