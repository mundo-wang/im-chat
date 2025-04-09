package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mundo-wang/wtool/wlog"
	"github.com/mundo-wang/wtool/wresp"
	"im-chat/code"
	"im-chat/service"
	"strconv"
	"strings"
)

type UserApi struct {
	service.UserService
	service.ContactService
}

func GetUserApi() *UserApi {
	return &UserApi{}
}

func (api *UserApi) CreateUser(c *gin.Context) (interface{}, error) {
	req := &service.CreateUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	// 不知道前端有没有做去掉首尾空格的处理，我这里先做一下
	userName := strings.TrimSpace(req.UserName)
	password := strings.TrimSpace(req.Password)
	rePassword := strings.TrimSpace(req.RePassword)
	if userName == "" || password == "" {
		return nil, code.UserNameOrPasswordEmpty
	}
	if password != rePassword {
		return nil, code.PasswordMismatch
	}
	err = api.UserService.CreateUser(userName, password)
	if err != nil {
		// 如果service里的代码返回了错误码，则需要进行这个判断，否则不需要
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.UserService.CreateUser failed").Err(err).Field("req", req).Log()
		}
		return nil, err
	}
	return nil, nil
}

func (api *UserApi) Login(c *gin.Context) (interface{}, error) {
	req := &service.FindByNamePwdReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Log()
		return nil, err
	}
	userName := strings.TrimSpace(req.UserName)
	password := strings.TrimSpace(req.Password)
	resp, jwtToken, err := api.UserService.FindByNamePwd(userName, password)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.UserService.FindByNamePwd failed").Err(err).Field("req", req).Log()
		}
		return nil, err
	}
	c.Header("Authorization", jwtToken)
	return resp, nil
}

func (api *UserApi) UpdateUser(c *gin.Context) (interface{}, error) {
	return nil, nil
}

func (api *UserApi) SearchFriends(c *gin.Context) (interface{}, error) {
	userIdStr := c.Query("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		wlog.Error("call strconv.Atoi failed").Err(err).Field("userId", userIdStr).Log()
		return nil, err
	}
	friends, err := api.UserService.SearchFriends(userId)
	if err != nil {
		wlog.Error("call api.UserService.SearchFriends failed").Err(err).Field("userId", userId).Log()
		return nil, err
	}
	return friends, nil
}

func (api *UserApi) ChangePassword(c *gin.Context) (interface{}, error) {
	req := &service.ChangePasswordReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		wlog.Error("call c.ShouldBindJSON failed").Err(err).Field("req", req).Log()
		return nil, err
	}
	if req.Password == req.NewPassword {
		return nil, code.PasswordUnchanged
	}
	err = api.UserService.ChangePassword(req.UserName, req.Password, req.NewPassword)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.UserService.ChangePassword failed").Err(err).Field("req", req).Log()
		}
		return nil, err
	}
	return nil, nil
}

func (api *UserApi) AddFriend(c *gin.Context) (interface{}, error) {
	userCode := c.Query("userCode")
	userId := c.GetInt("userId")
	err := api.UserService.AddFriend(userCode, userId)
	if err != nil {
		if !wresp.IsErrorCode(err) {
			wlog.Error("call api.UserService.AddFriend failed").Err(err).Field("userCode", userCode).Log()
		}
		return nil, err
	}
	return nil, nil
}
