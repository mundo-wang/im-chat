package db

import (
	"fmt"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"im-chat/conf"
)

func GetDB() *gorm.DB {
	userName := conf.Config.MySQL.UserName
	password := conf.Config.MySQL.Password
	ip := conf.Config.MySQL.IP
	port := conf.Config.MySQL.Port
	dbName := conf.Config.MySQL.DBName
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
