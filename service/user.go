package service

import (
	"errors"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gorm"
	"im-chat/code"
	"im-chat/dao/model"
	"im-chat/dao/query"
	"im-chat/utils"
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
		Name:     userName,
		Password: signature,
		Salt:     salt,
	}
	err = usersQ.Omit(usersQ.LoginTime, usersQ.HeartbeatTime, usersQ.LogoutTime).Create(user)
	if err != nil {
		wlog.Error("call usersQ.Create failed").Err(err).Field("user", user).Log()
		return err
	}
	return nil
}

func (u *UserService) FindByNamePwd(userName, password string) (*model.Users, error) {
	usersQ := query.Users
	user, err := usersQ.Where(usersQ.Name.Eq(userName)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, code.UserNameNotExist
		}
		wlog.Error("call usersQ.First failed").Err(err).Field("userName", userName).Log()
		return nil, err
	}
	check := utils.VerifySignature(password, user.Salt, user.Password)
	if !check {
		return nil, code.PasswordNotCorrect
	}
	return user, nil
}

func (u *UserService) SearchFriends(userId int) ([]*model.Users, error) {
	usersQ := query.Users
	contactsQ := query.Contacts
	friendIds := make([]int, 0)
	err := contactsQ.Select(contactsQ.TargetID).Where(contactsQ.OwnerID.Eq(userId),
		contactsQ.Type.Eq(utils.ContactTypeFriend)).Scan(&friendIds)
	if err != nil {
		wlog.Error("call contactsQ.Scan failed").Err(err).Field("userId", userId).Log()
		return nil, err
	}
	friends, err := usersQ.Where(usersQ.ID.In(friendIds...)).Find()
	if err != nil {
		wlog.Error("call usersQ.Find failed").Err(err).Field("friendIds", friendIds).Log()
		return nil, err
	}
	return friends, nil
}
