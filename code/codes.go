package code

import "github.com/mundo-wang/wtool/wresp"

var (
	UserNameOrPasswordEmpty = wresp.NewErrorCode(1001, "用户名或密码为空，请重新输入")
	PasswordMismatch        = wresp.NewErrorCode(1002, "两次密码输入不一致，请重新输入")
	UserNameAlreadyExist    = wresp.NewErrorCode(1003, "该用户名已经存在")
)
