package service

import (
	"errors"
	"github.com/jinzhu/copier"
	"github.com/mundo-wang/wtool/wlog"
	"gorm.io/gorm"
	"im-chat/code"
	"im-chat/dao/model"
	"im-chat/utils"
)

type UserService struct {
}

func (u *UserService) CreateUser(userName, password string) error {
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
	var userCode string
	for {
		userCode, err = utils.GenerateRandomDigits(10)
		if err != nil {
			wlog.Error("call utils.GenerateRandomDigits failed").Err(err).Log()
			return err
		}
		count, err = usersQ.Where(usersQ.UserCode.Eq(userCode)).Count()
		if err != nil {
			wlog.Error("call usersQ.Count failed").Err(err).Field("userCode", userCode).Log()
			return err
		}
		if count == 0 {
			break
		}
	}
	user := &model.Users{
		UserCode: userCode,
		Name:     userName,
		Password: signature,
		Salt:     salt,
	}
	err = usersQ.Omit(usersQ.LoginTime, usersQ.HeartbeatTime).Create(user)
	if err != nil {
		wlog.Error("call usersQ.Create failed").Err(err).Field("user", user).Log()
		return err
	}
	return nil
}

func (u *UserService) FindByNamePwd(userName, password string) (*FindByNamePwdResp, string, error) {
	user, err := usersQ.Where(usersQ.Name.Eq(userName)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", code.UserNameNotExist
		}
		wlog.Error("call usersQ.First failed").Err(err).Field("userName", userName).Log()
		return nil, "", err
	}
	check := utils.VerifySignature(password, user.Salt, user.Password)
	if !check {
		return nil, "", code.PasswordNotCorrect
	}
	jwtToken, err := utils.GenerateJwtToken(user.ID, user.Name, user.Phone)
	if err != nil {
		wlog.Error("call utils.GenerateJwtToken failed").Err(err).Field("userName", userName).Log()
		return nil, "", err
	}
	resp := &FindByNamePwdResp{}
	err = copier.Copy(resp, user)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("userName", userName).Log()
		return nil, "", err
	}
	return resp, jwtToken, nil
}

func (u *UserService) UpdateUser(req *UpdateUserReq) error {
	user := &model.Users{}
	err := copier.Copy(user, req)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("req", req).Log()
		return err
	}
	_, err = usersQ.Updates(user)
	if err != nil {
		wlog.Error("call usersQ.Updates failed").Err(err).Field("req", req).Log()
		return err
	}
	return nil
}

func (u *UserService) SearchFriends(userId int) ([]*SearchFriendsResp, error) {
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
	resp := []*SearchFriendsResp{}
	err = copier.Copy(&resp, friends)
	if err != nil {
		wlog.Error("call copier.Copy failed").Err(err).Field("friendIds", friendIds).Log()
		return nil, err
	}
	return resp, nil
}

func (u *UserService) ChangePassword(userName, password, newPassword string) error {
	user, err := usersQ.Where(usersQ.Name.Eq(userName)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.UserNameNotExist
		}
		wlog.Error("call usersQ.First failed").Err(err).Field("userName", userName).Log()
		return err
	}
	check := utils.VerifySignature(password, user.Salt, user.Password)
	if !check {
		return code.PasswordNotCorrect
	}
	signature, salt, err := utils.GenerateSignature(newPassword)
	if err != nil {
		wlog.Error("call utils.GenerateSignature failed").Err(err).Field("userName", userName).Log()
		return err
	}
	user.Password = signature
	user.Salt = salt
	_, err = usersQ.Where(usersQ.ID.Eq(user.ID)).Updates(user)
	if err != nil {
		wlog.Error("call usersQ.Create failed").Err(err).Log()
		return err
	}
	return nil
}

func (u *UserService) AddFriend(userCode string, userId int) error {
	targetUser, err := usersQ.Where(usersQ.UserCode.Eq(userCode)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.UserCodeNotExist
		}
		wlog.Error("call usersQ.First failed").Err(err).Field("userCode", userCode).Log()
		return err
	}
	count, err := contactsQ.Where(contactsQ.Where(contactsQ.OwnerID.Eq(userId), contactsQ.TargetID.Eq(targetUser.ID)).
		Or(contactsQ.OwnerID.Eq(targetUser.ID), contactsQ.TargetID.Eq(userId))).Where(contactsQ.Type.Eq(utils.ContactTypeFriend)).Count()
	if err != nil {
		wlog.Error("call contactsQ.Count failed").Err(err).Field("userId", userId).Field("targetId", targetUser.ID).Log()
		return err
	}
	if count != 0 {
		return code.AlreadyFriends
	}
	contact1 := &model.Contacts{
		OwnerID:  userId,
		TargetID: targetUser.ID,
		Type:     utils.ContactTypeFriend,
	}
	contact2 := &model.Contacts{
		OwnerID:  targetUser.ID,
		TargetID: userId,
		Type:     utils.ContactTypeFriend,
	}
	err = contactsQ.Create(contact1, contact2)
	if err != nil {
		wlog.Error("call contactsQ.Create failed").Err(err).Field("contact1", contact1).Log()
		return err
	}
	return nil
}
