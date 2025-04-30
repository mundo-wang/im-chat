package utils

const (
	ContactTypeFriend = 1 // 好友
	ContactTypeGroup  = 2 // 群组

	Authorization = "Authorization"

	GenerateStatusGenerating = 1 // 生成中
	GenerateStatusGenerated  = 2 // 已生成
	GenerateStatusFailed     = 3 // 生成失败
	GenerateStatusSubmitted  = 4 // 已提交

	QuestionStatusUnpublished = 0 // 题目未发布
	QuestionStatusPublished   = 1 // 题目已发布

	ContextUserIDKey   = "userId"
	ContextUserNameKey = "userName"
	ContextPhoneKey    = "phone"

	QuestionTypeChoice   = 0 // 选择题
	QuestionTypeJudgment = 1 // 判断题
)
