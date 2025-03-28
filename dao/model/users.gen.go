// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUsers = "users"

// Users mapped from table <users>
type Users struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Password      string         `gorm:"column:password;not null" json:"password"`
	Phone         string         `gorm:"column:phone;not null" json:"phone"`
	Email         string         `gorm:"column:email;not null" json:"email"`
	Identity      string         `gorm:"column:identity" json:"identity"`
	ClientInfo    string         `gorm:"column:client_info" json:"client_info"`
	LoginTime     time.Time      `gorm:"column:login_time" json:"login_time"`
	HeartbeatTime time.Time      `gorm:"column:heartbeat_time" json:"heartbeat_time"`
	LogoutTime    time.Time      `gorm:"column:logout_time" json:"logout_time"`
	IsLogout      int            `gorm:"column:is_logout;not null;comment:0为未登出，1为已登出" json:"is_logout"` // 0为未登出，1为已登出
	DeviceInfo    string         `gorm:"column:device_info" json:"device_info"`
	Avatar        string         `gorm:"column:avatar" json:"avatar"`
	Salt          string         `gorm:"column:salt" json:"salt"`
	CreatedAt     time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Users's table name
func (*Users) TableName() string {
	return TableNameUsers
}
