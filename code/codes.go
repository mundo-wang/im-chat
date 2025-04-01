package code

import "github.com/mundo-wang/wtool/wresp"

var (
	UserNameOrPasswordEmpty = wresp.NewErrorCode(1001, "用户名或密码为空，请重新输入")
	PasswordMismatch        = wresp.NewErrorCode(1002, "两次密码输入不一致，请重新输入")
	UserNameAlreadyExist    = wresp.NewErrorCode(1003, "该用户名已经存在")
	UserNameNotExist        = wresp.NewErrorCode(1004, "查询的用户名不存在")
	PasswordNotCorrect      = wresp.NewErrorCode(1005, "密码不正确，请重新输入")
	PasswordUnchanged       = wresp.NewErrorCode(1006, "新密码不能与旧密码相同，请重新输入")
	JwtTokenInvalid         = wresp.NewErrorCode(1007, "无效的身份令牌，请重新登录")
)
