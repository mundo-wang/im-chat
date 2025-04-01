package service

import "time"

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
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	LoginTime time.Time `json:"login_time"`
}

type ChangePasswordReq struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type SearchFriendsResp struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	LoginTime time.Time `json:"login_time"`
}
