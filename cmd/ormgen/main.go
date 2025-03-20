package main

import (
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gen"
	"im-chat/cmd/db"
	"im-chat/utils"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		wlog.Fatal("call utils.InitConfig failed").Err(err).Log()
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db.GetDB())
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
