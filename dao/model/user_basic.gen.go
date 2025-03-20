// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserBasic = "user_basic"

// UserBasic mapped from table <user_basic>
type UserBasic struct {
	ID            int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string         `gorm:"column:name" json:"name"`
	PassWord      string         `gorm:"column:pass_word" json:"pass_word"`
	Phone         string         `gorm:"column:phone" json:"phone"`
	Email         string         `gorm:"column:email" json:"email"`
	Identity      string         `gorm:"column:identity" json:"identity"`
	ClientIP      string         `gorm:"column:client_ip" json:"client_ip"`
	ClientPort    string         `gorm:"column:client_port" json:"client_port"`
	LoginTime     time.Time      `gorm:"column:login_time" json:"login_time"`
	HeartbeatTime time.Time      `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	LoginOutTime  time.Time      `gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool           `gorm:"column:is_logout" json:"is_logout"`
	DeviceInfo    string         `gorm:"column:device_info" json:"device_info"`
	Salt          string         `gorm:"column:salt" json:"salt"`
	Avatar        string         `gorm:"column:avatar" json:"avatar"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserBasic's table name
func (*UserBasic) TableName() string {
	return TableNameUserBasic
}
