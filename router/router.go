package router

import (
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/api"
)

func SetRouter(s *wresp.Server) {
	r := s.Router
	r.MaxMultipartMemory = 8 << 20
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
		user.POST("/login", s.WrapHandler(api.GetUserApi().Login))
		user.Use(s.WrapMiddleware(api.CheckAuthorization)) // 上面两个用户接口不走鉴权中间件
		user.POST("/update", s.WrapHandler(api.GetUserApi().UpdateUser))
		user.GET("/searchFriends", s.WrapHandler(api.GetUserApi().SearchFriends))
		user.POST("/changePassword", s.WrapHandler(api.GetUserApi().ChangePassword))
		user.GET("/addFriend", s.WrapHandler(api.GetUserApi().AddFriend))
	}

	// 与群组有关的接口
	community := r.Group("/community", s.WrapMiddleware(api.CheckAuthorization))
	{
		community.GET("/loadByUserId", s.WrapHandler(api.GetCommunityApi().LoadByUserId))
		community.POST("/create", s.WrapHandler(api.GetCommunityApi().Create))
		community.GET("/joinGroup", s.WrapHandler(api.GetCommunityApi().JoinGroup))
	}

	// 上传下载文件接口
	attach := r.Group("/attach", s.WrapMiddleware(api.CheckAuthorization))
	{
		attach.POST("/upload", s.WrapHandler(api.Upload))
		attach.GET("/download", s.WrapFileDownload(api.Download, false)) // false代表不触发下载
	}

	// 试题有关接口
	question := r.Group("/question", s.WrapMiddleware(api.CheckAuthorization))
	{
		question.GET("/generate", s.WrapStreamHandler(api.GetQuestionSessionApi().GenerateQuestions))
		question.GET("/checkUnpublishedSession", s.WrapHandler(api.GetQuestionSessionApi().CheckUnpublishedSession))
		question.GET("/getSessionQuestions", s.WrapHandler(api.GetQuestionSessionApi().GetSessionQuestions))
	}
}
