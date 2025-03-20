package main

import (
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/cmd/db"
	"im-chat/dao/query"
	"im-chat/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		wlog.Fatal("call utils.InitConfig failed").Err(err).Log()
	}
	query.SetDefault(db.GetDB())
}
