package main

import (
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/cmd/db"
	"im-chat/dao/query"
	"im-chat/router"
	"im-chat/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		wlog.Error("call utils.InitConfig failed").Err(err).Log()
		return
	}
	query.SetDefault(db.GetDB())
	wlog.Info("InitMySQL complete").Log()
	utils.InitRedis()
	r := router.SetRouter()
	err = r.Run(fmt.Sprintf(":%d", utils.Config.Server.Port))
	if err != nil {
		wlog.Error("call r.Run failed").Err(err).Field("serverPort", utils.Config.Server.Port).Log()
		return
	}
}
