package api

import (
	"github.com/gin-gonic/gin"
	"im-chat/code"
	"im-chat/service"
	"strings"
)

type UserApi struct {
	service.UserService
}

func GetUserApi() *UserApi {
	return &UserApi{}
}

func (u *UserApi) CreateUser(c *gin.Context) (interface{}, error) {
	req := &service.CreateUserReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
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
	err = u.UserService.CreateUser(userName, password, rePassword)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
