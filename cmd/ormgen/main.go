package main

import (
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gen"
	"gorm.io/gorm"
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
	customTypeMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"int": func(columnType gorm.ColumnType) string {
			return "int"
		},
		"tinyint": func(columnType gorm.ColumnType) string {
			return "int"
		},
	}
	g.WithDataTypeMap(customTypeMap)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
