package router

import (
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/api"
)

func SetRouter(s *wresp.Server) {
	r := s.Router
	// 将asset/目录映射到/asset接口路径，提供静态文件访问功能
	r.Static("/asset", "asset/")
	// 将/favicon.ico路由直接映射到asset/images/favicon.ico这个具体的文件，用于提供单个静态文件的访问
	r.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	// 加载views目录下所有子目录的HTML模板，用于Gin的HTML渲染
	r.LoadHTMLGlob("views/**/*")

	// 与渲染前端页面有关的接口
	front := r.Group("/front")
	{
		front.GET("/index", s.WrapHandler(api.GetFrontApi().GetIndex))
		front.GET("/toRegister", s.WrapHandler(api.GetFrontApi().ToRegister))
		front.GET("/toChat", s.WrapHandler(api.GetFrontApi().ToChat))
	}
	// 与用户相关的接口
	user := r.Group("/user")
	{
		user.POST("/create", s.WrapHandler(api.GetUserApi().CreateUser))
		user.POST("/findByNamePwd", s.WrapHandler(api.GetUserApi().FindByNamePwd))
	}
}
