package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/cmd/db"
	"im-chat/dao/query"
	"im-chat/router"
	"im-chat/service"
	"im-chat/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		wlog.Error("call utils.InitConfig failed").Err(err).Log()
		return
	}
	query.SetDefault(db.GetDB())
	service.InitDao()
	wlog.Info("InitMySQL complete!").Log()
	s := NewServer()
	err = s.Router.Run(fmt.Sprintf(":%d", utils.Config.Server.Port))
	if err != nil {
		wlog.Error("call r.Run failed").Err(err).Field("serverPort", utils.Config.Server.Port).Log()
		return
	}
}

func NewServer() *wresp.Server {
	s := &wresp.Server{
		Router: gin.Default(),
	}
	router.SetRouter(s)
	return s
}
