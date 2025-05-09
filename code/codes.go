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
	UserIdNotSet            = wresp.NewErrorCode(1008, "用户ID未获取到")
	UserIdFormatWrong       = wresp.NewErrorCode(1009, "用户ID的格式有误")
	GroupNameAlreadyExist   = wresp.NewErrorCode(1010, "该群名已经存在")
	UserCodeNotExist        = wresp.NewErrorCode(1011, "查询的用户code不存在")
	AlreadyFriends          = wresp.NewErrorCode(1012, "你们已经是好友了")
	GroupCodeNotExist       = wresp.NewErrorCode(1013, "查询的群聊code不存在")
	AlreadyInGroup          = wresp.NewErrorCode(1014, "你已经在此群里了")
	FileNotExist            = wresp.NewErrorCode(1015, "文件未找到")
	JwtTokenNotExist        = wresp.NewErrorCode(1016, "未获取到Token")
	AnotherPerson           = wresp.NewErrorCode(1017, "不能修改他人的个人信息")
	QuestionGenerating      = wresp.NewErrorCode(1018, "该批次问题仍在生成中")
)
