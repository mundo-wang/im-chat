package service

import (
	"time"
)

type CreateUserReq struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}

type FindByNamePwdReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type FindByNamePwdResp struct {
	ID        int       `json:"id"`
	UserCode  string    `json:"userCode"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	LoginTime time.Time `json:"loginTime"`
}

type ChangePasswordReq struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type UpdateUserReq struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type SearchFriendsResp struct {
	ID       int    `json:"id"`
	UserCode string `json:"userCode"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type LoadByUserIdResp struct {
	ID            int    `json:"id"`
	CommunityCode string `json:"communityCode"`
	Name          string `json:"name"`
	OwnerID       int    `json:"ownerID"`
	Avatar        string `json:"avatar"`
	Type          int    `json:"type"` // 0.默认 1.兴趣爱好 2.行业交流 3.生活休闲
	Description   string `json:"description"`
}

type CreateCommunityReq struct {
	Name        string `json:"name"`
	OwnerID     int    `json:"ownerID"`
	Avatar      string `json:"avatar"`
	Type        int    `json:"type"`
	Description string `json:"description"`
}

type GenerateQuestionsReq struct {
	PositionId int    `form:"positionId"` // 岗位id
	AgentCode  string `form:"agentCode"`  // 智能体id
	Type       int    `form:"type"`       // 题型（0=选择题，1=判断题）
	Number     int    `form:"number"`     // 题目数量
}

type GenerateQuestionsResp struct {
	ID      int       `json:"id"`      // 题目id
	Title   string    `json:"title"`   // 题干内容
	Answer  string    `json:"answer"`  // 正确答案（如：A、B、C、D）
	Options []Options `json:"options"` // 选项列表（选择题时必填，判断题可为空）
}

type Options struct {
	OptionKey  string `json:"optionKey"`  // 选项标识（如：A、B、C、D）
	OptionText string `json:"optionText"` // 选项内容
}

type GetSessionQuestionsResp struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Type       int       `json:"type"`
	PositionId int       `json:"positionId"`
	AgentCode  string    `json:"agentCode"`
	Answer     string    `json:"answer"`
	Options    []Options `json:"options"`
}
