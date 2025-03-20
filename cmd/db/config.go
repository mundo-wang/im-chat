package db

import (
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"im-chat/utils"
)

func GetDB() *gorm.DB {
	userName := utils.Config.MySQL.UserName
	password := utils.Config.MySQL.Password
	ip := utils.Config.MySQL.IP
	port := utils.Config.MySQL.Port
	dbName := utils.Config.MySQL.DBName
	dsnFmt := "%v:%v@(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFmt, userName, password, ip, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		wlog.Fatal("call gorm.Open failed").Err(err).Log()
	}
	return db
}
