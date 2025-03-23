package utils

import "github.com/mundo-wang/wtool/wresp"

var (
	UserNotFound     = wresp.NewErrorCode(10008, "未找到对应用户，请检查用户是否存在")
	CreateUserFailed = wresp.NewErrorCode(10009, "创建用户时出错，请检查创建参数")
)
