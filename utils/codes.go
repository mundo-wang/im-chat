package utils

import "github.com/mundo-wang/wtool/wresp"

var (
	// 错误码创建示例
	InvalidInput = wresp.NewErrorCode(10001, "提交的数据格式无效，请检查输入的内容")
	Unauthorized = wresp.NewErrorCode(10002, "未登录或权限不足，无法访问此资源")
	Forbidden    = wresp.NewErrorCode(10003, "访问被拒绝，您没有权限操作此资源，请联系管理员")
	NotFound     = wresp.NewErrorCode(10004, "请求的资源未找到，请确认URL是否正确")
	Timeout      = wresp.NewErrorCode(10005, "请求超时，请稍后重试")
)
