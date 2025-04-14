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
