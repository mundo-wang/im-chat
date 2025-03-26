package service

import (
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/code"
	"im-chat/dao/model"
	"im-chat/dao/query"
	"im-chat/utils"
	"time"
)

type UserService struct {
}

func (u *UserService) CreateUser(userName, password string) error {
	usersQ := query.Users
	count, err := usersQ.Where(usersQ.Name.Eq(userName)).Count()
	if err != nil {
		wlog.Error("call usersQ.Count() failed").Err(err).Field("userName", userName).Log()
		return err
	}
	if count != 0 {
		return code.UserNameAlreadyExist
	}
	signature, salt, err := utils.GenerateSignature(password)
	if err != nil {
		wlog.Error("call utils.GenerateSignature failed").Err(err).Log()
		return err
	}
	user := &model.Users{
		Name:          userName,
		Password:      signature,
		Salt:          salt,
		LoginTime:     time.Now(), // 这三条先为了避免数据库报错加在这里，后续看看是调整表还是怎么样
		HeartbeatTime: time.Now(),
		LogoutTime:    time.Now(),
	}
	err = usersQ.Create(user)
	if err != nil {
		wlog.Error("call usersQ.Create failed").Err(err).Field("user", user).Log()
		return err
	}
	return nil
}
