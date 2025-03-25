package service

import (
	"github.com/mundo-wang/wtool/wlog"
	"im-chat/code"
	"im-chat/dao/query"
)

type UserService struct {
}

func (u *UserService) CreateUser(userName, password, rePassword string) error {
	usersQ := query.Users
	count, err := usersQ.Where(usersQ.Name.Eq(userName)).Count()
	if err != nil {
		wlog.Error("call usersQ.Count() failed").Err(err).Field("userName", userName).Log()
		return err
	}
	if count != 0 {
		return code.UserNameAlreadyExist
	}

	return nil
}
